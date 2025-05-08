package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/mousavisam/streaming-event-processor/pkg/models"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) (*Producer, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		Balancer:     &kafka.Hash{}, // ensures same key goes to same partition
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
		Key:   []byte(event.SourceId), // preserves order per source
		Value: payload,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
