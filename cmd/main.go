package main

import (
	"log"
	"github.com/izzatbey/soc-kafka-preprocess/internal/server"
)

func main() {
	log.Println("🚀 Starting Go Preprocess Service on port 8081...")
	server.Start()
}
