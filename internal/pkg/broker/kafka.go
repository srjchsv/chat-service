package broker

import (
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitProducer(host string, timeout time.Duration) (*kafka.Producer, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": host,
		})
		if err == nil {
			return p, err
		}
		fmt.Printf("Failed to initialize Kafka producer: %v. Retrying...\n", err)
		time.Sleep(1 * time.Second)
	}

	return nil, errors.New("timed out waiting for Kafka producer to initialize")
}
