package publisher

import (
	"testing"
	publisher "Ponderada2/publisher"
)

func TestConnectToMQTT(t *testing.T) {
	clientID := "test_client"
	client := publisher.ConnectToMQTT(clientID)

	if client == nil {
		t.Errorf("Expected a valid MQTT client, but got nil")
	}

	if client.IsConnected() != true {
		t.Errorf("Expected MQTT client to be connected, but it's not")
	}
}
