package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mousavisam/streaming-event-processor/internal/ingest/http"
	"github.com/mousavisam/streaming-event-processor/internal/kafka"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, falling back to system environment")
	}

	fmt.Println("Loaded brokers from env:", os.Getenv("KAFKA_BROKERS"))

	kafkaProducer, err := kafka.NewProducer()
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	router := gin.Default()
	httpHandler := http.NewHandler(kafkaProducer)
	httpHandler.RegisterRoutes(router)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // fallback default
	}

	router.Run(":" + port)
}
