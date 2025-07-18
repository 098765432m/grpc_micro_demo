package kafka

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafkaWriter(brokers []string, topic string) {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func ProduceHotelCreatedEvent(event any) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(time.Now().String()),
		Value: data,
	}

	return kafkaWriter.WriteMessages(context.Background(), msg)
}
