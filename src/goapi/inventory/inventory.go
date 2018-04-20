package main

import (
	// "encoding/json"
	"fmt"
	// "net/http"
	"github.com/go-redis/redis"
	// "github.com/gorilla/mux"
	// "github.com/satori/go.uuid"
	// "github.com/unrolled/render"
)

func NewClient() {
	client := redis.NewClient(&redis.Options{
		Addr: 		"localhost:6379",
		Password: 	"",
		DB: 		1,
		})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}



