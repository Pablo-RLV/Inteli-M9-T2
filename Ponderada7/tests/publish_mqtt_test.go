package tests

import (
	common "Ponderada7/common"
	"testing"
)

func TestPublishToBroker(t *testing.T) {
	client := common.ConnectToMQTT("test_client", nil)
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
