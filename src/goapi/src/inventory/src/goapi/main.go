package main

import (
	"os"
)

func main() {
	server := NewServer()
	server.Init()
	server.Run()
}