package main

import (
	"testing"
)

func TestPublishToBroker(t *testing.T) {
	client := ConnectToMQTT("test_client")
	topic := "test_topic"
	msg := "test_message"
	token := client.Publish(topic, 0, false, msg)
	token.Wait()

	if token.Error() != nil {
		t.Errorf("Error publishing message: %s", token.Error())
	}
	if token.Wait() && token.Error() != nil {
		t.Errorf("Error publishing message: %s", token.Error())
	}
}