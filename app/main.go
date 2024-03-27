package main

import (
	"encoding/json"
	"fmt"

	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/data"
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/pubSub"
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/pubSub/consumer"
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/pubSub/producer"
	"github.com/John-Dembaremba/KAFKA-PRACTICE.git/schema"
)

func main() {

	kafkaHost := "localhost:9092" // Use the Kafka container's hostname
	connData := schema.KafkaConn{
		Address:    kafkaHost,
		Topic:      "person",
		Partitions: 1,
	}

	fmt.Println("Producer connecting...")
	prod, err := pubSub.ProducerConn(&connData)
	if err != nil {
		fmt.Println("Producer connection failed.")
		prod.Close()
	}

	fmt.Println("Producer connected successfully")
	dataHandler := data.New()
	for i := 1; i <= 3; i++ {

		person := dataHandler.GetPersonData()
		msg, err := json.Marshal(&person)
		if err != nil {
			fmt.Println("Failed to serializer message")
			break
		}

		e := producer.PublishMsg(prod, msg, connData.Topic, int32(connData.Partitions))
		if e != nil {
			fmt.Printf("Producer failed publish msg with error: %e\n", err)
			break
		}
		fmt.Printf("%d] published\n", i)
	}
	prod.Flush(15 * 1000)
	fmt.Println("Producer connection closing...")
	prod.Close()

	fmt.Println("Consumer connecting...")
	cons, err := pubSub.ConsumerConn(&connData)
	if err != nil {
		fmt.Printf("Consumer connection failed with error: %e.\n", err)
		cons.Close()
	}

	fmt.Println("Consumer connected successfully")
	bytesMgs, err := consumer.ConsumeMsg(cons, connData.Topic)
	if err != nil {
		fmt.Println("Consumer failed consuming message maybe its timeout")
		cons.Close()
	}
	var person schema.Person

	if err := json.Unmarshal(bytesMgs, &person); err != nil {
		fmt.Printf("Consumer process for deserialization failed with error: %e\n", err)
		cons.Close()
		return
	}
	fmt.Println("Reading data....")
	fmt.Println(person)
	fmt.Println("Consumer connection closing...")
	cons.Close()

}
