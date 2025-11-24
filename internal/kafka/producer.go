package kafka

import (
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer struct {
	p *kafka.Producer
}

func NewProducerWithBroker(broker string) *Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": 		broker,
		"enable.idempotence":       true,
		"linger.ms":                10,
		"batch.size":               65536,
		"acks":                     "all",
		"message.send.max.retries": 5,
		"retry.backoff.ms":         100,
		"socket.keepalive.enable":  true,
		"go.batch.producer":        true,

	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}

	go func() {
        for e := range p.Events() {
            switch ev := e.(type) {
            case *kafka.Message:
                if ev.TopicPartition.Error != nil {
                    log.Printf("⚠️ Delivery failed: %v", ev.TopicPartition.Error)
                }
            }
        }
    }()

	return &Producer{p: p}
}

func (kp *Producer) Publish(topic string, value []byte) {
    for {
        err := kp.p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value:          value,
        }, nil)

        if err == nil {
            return
        }

        // Queue full: wait a bit and retry
        if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrQueueFull {
            log.Printf("⚠️ Local queue full, backing off...")
            time.Sleep(100 * time.Millisecond)
            continue
        }

        log.Printf("⚠️ Failed to publish message: %v", err)
        return
    }
}

func (kp *Producer) Close() {
	kp.p.Flush(1000)
	kp.p.Close()
}
