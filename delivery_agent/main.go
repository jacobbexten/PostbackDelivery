package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
	//"regexp"
	//"net/url"
)

// connecting to Redis server
func client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return client
}

// creates log.txt (if not already existent) to output responses to
func logger() *os.File {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	return file
}

type Postback struct {
	Endpoint struct {
		Method string `json:"method"`
		URL    string `json:"url"`
	} `json:"endpoint"`
	Data map[string]string `json:"data"`
}

// gets object from the Redis queue
func getFromRedis(client *redis.Client, data string) (*Postback, error) {
	var postback Postback
	result, err := client.BRPop(0, data).Result()

	if err != nil {
		log.Println("Error pulling from queue")
	} else {
		json.Unmarshal([]byte(result[1]), &postback)
	}
	log.Println("Preprocessed from Redis: ", result)
	return &postback, nil
}

type delivery struct {
	deliveryTime string
	responseCode string
	responseTime string
	responseBody string
}

// reformat the URL from the JSON to use as GET request
func reformatURL(data Postback) string {
	for key, value := range data.Data {
		fmt.Println("REACHED")
		value = url.QueryEscape(value)
		re := regexp.MustCompile(`\{` + key + `\}`)
		data.Endpoint.URL = re.ReplaceAllString(data.Endpoint.URL, value)
	}

	//fmt.Println(re.ReplaceAllString(data.Endpoint.URL,data.Data[0].Mascot))
	fmt.Println(data.Data)

	return data.Endpoint.URL
}

// sends GET requests to endpoint
func sendRequest(URL string) (*delivery, error) {
	time_start := time.Now()

	var delivery_data delivery
	response, err := http.Get(URL)

	time_end := time.Now()

	delivery_data.deliveryTime = time_end.String()
	delivery_data.responseTime = time_end.Sub(time_start).String()

	if err != nil {
		log.Fatal(err)
	} else {
		defer response.Body.Close()
		delivery_data.responseCode = string(response.StatusCode)
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		body_string := string(body)
		log.Println(body_string)
	}
	return &delivery_data, nil
}

func main() {
	client := client()
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to redis %v", err)
	}
	fmt.Println(pong, err)

	logger := logger()
	defer logger.Close()

	for {
		data, err := getFromRedis(client, "data")
		if err != nil {
			log.Println("Can't retrieve from Redis")
		}
		fmt.Println("HELLO")
		URL := reformatURL(*data)
		delivered, err := sendRequest(URL)

		if err != nil {
			log.Println("Error sending request")
		} else {
			log.Println("Delivery Time: ", delivered.deliveryTime)
			log.Println("Response Code: ", delivered.responseCode)
			log.Println("Response Time: ", delivered.responseTime)
			log.Println("Response Body: ", delivered.responseBody)

		}
	}
}
