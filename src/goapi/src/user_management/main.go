package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: ./user_management <config_file>\n")
		return
	}

	server := NewServer(os.Args[1])
	server.Init()
	server.Run()
}
