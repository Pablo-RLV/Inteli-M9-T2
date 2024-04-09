package common

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func ConnectToMQTT(clientID string, messagePubHandler MQTT.MessageHandler) MQTT.Client {
	LoadEnv()
	opts := SetOptions(clientID)
	if messagePubHandler != nil {
		opts.SetDefaultPublishHandler(messagePubHandler)
	}
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}