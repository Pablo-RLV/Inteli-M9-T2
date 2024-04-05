package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

func TestStorage(t *testing.T) {
	var bootstrap_servers = "localhost:29092,localhost:39092"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_servers,
		"client.id":         "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()
}