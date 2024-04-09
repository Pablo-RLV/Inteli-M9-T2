package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	common "Ponderada7/common"
	"log"
	"os"
)

func Kafka() {
	common.LoadEnv()
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("CONFLUENT_BOOTSTRAP_SERVER_SASL"),
		"security.protocol":  "SASL_SSL",
		"sasl.mechanisms":    "PLAIN",
		"sasl.username":      os.Getenv("CLUSTER_API_KEY"),
		"sasl.password":      os.Getenv("CLUSTER_API_SECRET"),
		"session.timeout.ms": 45000,
		"group.id":           "go-group-1",
		"auto.offset.reset":  "latest",
	}
	consumer, err := kafka.NewConsumer(configMap)
	if err != nil {
		log.Printf("Error creating kafka consumer: %v", err)
	}
	consumer.SubscribeTopics([]string{"sensor"}, nil)
	run := true
	for run {
		e := consumer.Poll(1000)
		switch ev := e.(type) {
		case *kafka.Message:
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", ev)
			run = false
		}
	}
	consumer.Close()
}