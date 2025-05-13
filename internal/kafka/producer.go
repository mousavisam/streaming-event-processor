package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mousavisam/streaming-event-processor/pkg/models"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer() (*Producer, error) {
	brokerStr := os.Getenv("KAFKA_BROKERS")
	if brokerStr == "" {
		return nil, fmt.Errorf("KAFKA_BROKERS env not set")
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		return nil, fmt.Errorf("KAFKA_TOPIC env not set")
	}

	fmt.Println("📌 KAFKA_BROKERS =", os.Getenv("KAFKA_BROKERS"))

	brokers := strings.Split(brokerStr, ",")

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Balancer:     &kafka.Hash{},
		RequiredAcks: int(kafka.RequireAll),
	})

	return &Producer{writer: writer}, nil
}

func (p *Producer) SendEvent(ctx context.Context, event *models.Event) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(event.SourceId),
		Value: payload,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
