package tests

import (
	"os"
	"testing"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	common "Ponderada7/common"
)

func TestKafkaCreation(t *testing.T){

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
		t.Errorf("Error creating kafka consumer: %v", err)
	}
	consumer.SubscribeTopics([]string{"sensor"}, nil)
}