package main

import (
	"log"
	"preprocess-service/internal/server"
)

func main() {
	log.Println("🚀 Starting Go Preprocess Service on port 8081...")
	server.Start()
}
