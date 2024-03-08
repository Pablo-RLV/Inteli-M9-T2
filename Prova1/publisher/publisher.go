package publisher

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

func ConnectToMQTT(clientID string) MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID(clientID)
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func PublishToBroker(client MQTT.Client, topic string, msg string) {
	token := client.Publish(topic, 1, false, msg)
	token.Wait()
}

func TimeToPublish() {
	time.Sleep(2 * time.Second)
}


func MessageSender(clientID string, tipo string) map[string]interface{} {
	var temperatura string
	if tipo == "freezer" {
		temperatura = RandFreezer(25, 15)
	} else {
		temperatura = RandGeladeira(10, 2)
	}
	data := map[string]interface{}{
		"id":          clientID,
		"tipo":        tipo,
		"temperatura": temperatura,
		"timestamp":  time.Now(),
	}
	return data
}

func Pub(id string, tipo string) {
	client := ConnectToMQTT(id)
	data := MessageSender(id, tipo)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting data to JSON", err)
		return
	}
	msg := string(jsonData)
	PublishToBroker(client, "/sensor", msg)
	// fmt.Println("Published:", msg)
	TimeToPublish()
}
