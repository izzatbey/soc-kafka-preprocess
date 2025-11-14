package kafka

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	p *kafka.Producer
}

func NewProducer() *Producer {
	broker := os.Getenv("KAFKA_BROKER")
	return NewProducerWithBroker(broker)
}

func NewProducerWithBroker(broker string) *Producer {
	if broker == "" {
		broker = "localhost:9092"
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	return &Producer{p: p}
}

func (kp *Producer) Publish(topic string, value []byte) {
	err := kp.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)
	if err != nil {
		log.Printf("⚠️ Failed to publish message: %v", err)
	}
}

func (kp *Producer) Close() {
	kp.p.Flush(1000)
	kp.p.Close()
}
