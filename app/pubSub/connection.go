package pubSub

import (
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/schema"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ProducerConn(connData *schema.KafkaConn) (*kafka.Producer, error) {
	return kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": connData.Address})
}

func ConsumerConn(connData *schema.KafkaConn) (*kafka.Consumer, error) {
	return kafka.NewConsumer(&kafka.ConfigMap{"bootstrap.servers": connData.Address,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest"})
}
