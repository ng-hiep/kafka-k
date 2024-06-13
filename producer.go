package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092"})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "myTopic"
	for _, word := range []string{"Hello", "Kafka", "from", "Golang", "OK"} {
		fmt.Println(word)

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic,
			Partition: kafka.PartitionAny},
			Value: []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
