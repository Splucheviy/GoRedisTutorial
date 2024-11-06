package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Hello world!")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ping)

	type Person struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Age        int    `json:"age"`
		Occupation string `json:"occupation"`
	}

	elliotID := uuid.NewString()

	jsonString, err := json.Marshal(Person{
		ID:         elliotID,
		Name:       "Elliot",
		Age:        30,
		Occupation: "Software Engineer",
	})

	if err != nil {
		fmt.Printf("Failed to marshal JSON: %s", err.Error())
		return
	}

	elliotKey := fmt.Sprintf("person:%s", elliotID)
	err = client.Set(context.Background(), elliotKey, jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
		return
	}

	val, err := client.Get(context.Background(), elliotKey).Result()
	if err != nil {
		fmt.Printf("Failed to get value from redis: %s", err.Error())
		return
	}

	fmt.Printf("value retrieved from redis: %s\n", val)
}
