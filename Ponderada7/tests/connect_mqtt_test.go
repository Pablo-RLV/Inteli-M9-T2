package tests

import (
	common "Ponderada7/common"
	"testing"
)

func TestConnectToMQTT(t *testing.T) {
	clientID := "test_client"
	client := common.ConnectToMQTT(clientID, nil)

	if client == nil {
		t.Errorf("Expected a valid MQTT client, but got nil")
	}

	if client.IsConnected() != true {
		t.Errorf("Expected MQTT client to be connected, but it's not")
	}
}
