package main

import (
	"log"
	"os"
	"time"

	"github.com/srjchsv/chat-service/internal/app/routes"
	"github.com/srjchsv/chat-service/internal/pkg/broker"
	"github.com/srjchsv/chat-service/internal/pkg/database"
)

func main() {
	db, err := database.InitDB(5, time.Second)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Initialize the Kafka producer
	producer, err := broker.InitProducer(os.Getenv("KAFKA_HOST"))
	if err != nil {
		log.Fatalf("Failed to initialize Kafka producer: %v", err)
	}

	defer producer.Close()
	router := routes.SetupRouter(db, producer)
	log.Fatal(router.Run(":8000"))
}