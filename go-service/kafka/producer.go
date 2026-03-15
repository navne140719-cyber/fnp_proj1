package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer() {
	writer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "order-created",
		Balancer: &kafka.LeastBytes{},
	}
	fmt.Println("Kafka Producer Initialized")
}

func PublishOrderEvent(data interface{}) error {
	jsonData, _ := json.Marshal(data)
	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: jsonData,
		},
	)
}
