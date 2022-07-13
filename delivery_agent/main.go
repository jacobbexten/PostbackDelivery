package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
)

// connecting to redis server
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

type postback struct {
	method   string `json:"method"`
	url      string `json:"url"`
	mascot   string `json:"mascot"`
	location string `json:"location"`
}

// gets object from the redis queue
func getFromRedis(client *redis.Client, data string) (*postback, error) {
	result, err := client.BRPop(0, data).Result()
	postback := postback{}
	if err != nil {
		fmt.Println(err)
	} else {
		err := json.Unmarshal([]byte(result[1]), &postback)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(result)
	return &postback, nil
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

	getFromRedis(client, "data")

}
