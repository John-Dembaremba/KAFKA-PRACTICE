package consumer

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func ConsumeMsg(conn *kafka.Consumer, topic string) ([]byte, error) {
	conn.Subscribe(topic, nil)
	msg, err := conn.ReadMessage(120 * time.Second) //Read msg for 2 mins.
	return msg.Value, err
}
