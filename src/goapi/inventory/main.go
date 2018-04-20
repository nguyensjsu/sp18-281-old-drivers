package main

import (
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "6379"
	}

	client := NewClient()
	client.Run(":" + port)
}