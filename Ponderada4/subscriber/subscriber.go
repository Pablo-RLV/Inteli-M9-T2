package subscriber

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	common "Ponderada4/common"
)

func GetMessagePubHandler() MQTT.MessageHandler {
	return func(client MQTT.Client, msg MQTT.Message) {
		payloadString := string(msg.Payload())
		var data map[string]interface{}
		err := json.Unmarshal([]byte(payloadString), &data)
		if err != nil {
			fmt.Println("Error parsing JSON payload:", err)
			return
		}
		idValue, found := data["ID"].(string)
		if !found {
			fmt.Println("ID not found in JSON payload")
			return
		}
		delete(data, "ID")
		newJSON, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}
		fmt.Printf("Received': %s from: %s, topic: %s \n", string(newJSON), idValue, msg.Topic())
	}
}

func ConnectToMQTT(clientID string) MQTT.Client {
	common.LoadEnv()
	opts := common.SetOptions(clientID)
	opts.SetDefaultPublishHandler(GetMessagePubHandler())
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func Sub() {
	id := "subscriber"
	client := ConnectToMQTT(id)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	if token := client.Subscribe("/sensor", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
}
