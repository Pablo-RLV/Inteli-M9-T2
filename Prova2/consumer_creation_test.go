package main

import (
	"testing"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestConsumerCreation(t *testing.T) {
	var bootstrap_servers = "localhost:29092,localhost:39092"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_servers,
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
}
