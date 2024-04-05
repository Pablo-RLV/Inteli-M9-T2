package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
	"testing"
	"math/rand"
)

func create_random_number() int{
	var random_number = rand.Intn(100)
	return random_number
}

func TestProducerCreation(t *testing.T) {
	var bootstrap_servers = "localhost:29092,localhost:39092"
	
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_servers,
		"client.id":         "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	topic := "qualidadeAr"
	data := map[string]interface{}{
		"idSensor":     "sensor_001",
		"timestamp":    time.Now(),
		"tipoPoluente": "PM2.5",
		"nivel":        create_random_number(),
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting data to JSON", err)
		return
	}
	msg := string(jsonData)


	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(fmt.Sprintf("%v", msg)),
	}, nil)
	producer.Flush(15 * 1000)

	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(fmt.Sprintf("%v", msg)),
	}, nil)
	producer.Flush(15 * 1000)

	contador := 0
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrap_servers,
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	consumer.SubscribeTopics([]string{topic}, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
			contador++
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	if contador != 2 {
		t.Errorf("Expected 2 messages, got %d", contador)
	}
}
