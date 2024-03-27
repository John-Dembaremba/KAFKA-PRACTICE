package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func PublishMsg(conn *kafka.Producer, message []byte, topic string, partition int32) error {
	m := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: partition,
		},
		Value: message,
	}
	return conn.Produce(&m, nil)
}
