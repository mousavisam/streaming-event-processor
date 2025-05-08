package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mousavisam/streaming-event-processor/internal/ingest/http"
	"github.com/mousavisam/streaming-event-processor/internal/kafka"
)

func main() {
	kafkaProducer, err := kafka.NewProducer([]string{"localhost:9092"}, "events")
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	router := gin.Default()
	httpHandler := http.NewHandler(kafkaProducer)
	httpHandler.RegisterRoutes(router)

	router.Run(":8080")
}
