package main

import (
	"testing"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func TestProducerCreation(t *testing.T) {
	var bootstrap_servers = "localhost:29092,localhost:39092"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_servers,
		"client.id":         "go-producer",
	})
	print(err)
	
	if err != nil {
		t.Fail()
		// panic(err)
	}
	defer producer.Close()
}
