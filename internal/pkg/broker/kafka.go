package broker

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func InitProducer(host string) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": host,
	})
	if err != nil {
		return p, err
	}

	return p, nil
}
