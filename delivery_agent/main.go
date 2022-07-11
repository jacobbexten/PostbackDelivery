package main

import (
	"fmt"
	"github.com/go-redis/redis"
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

func main() {
	fmt.Println("Hello World")

	client := client()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
