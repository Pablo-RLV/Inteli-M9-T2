package publisher

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"strconv"
	"time"
	common "Ponderada4/common"
)

func ConnectToMQTT(clientID string) MQTT.Client {
	common.LoadEnv()
	opts := common.SetOptions(clientID)
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

func MessageSender(clientID string) map[string]interface{} {
	data := map[string]interface{}{
		"ID":  clientID,
		"CO":  strconv.Itoa(rand.Intn(1000)) + " ppm",
		"NO2": strconv.Itoa(rand.Intn(10)) + " ppm",
		"NH3": strconv.Itoa(rand.Intn(500)) + " ppm",
	}
	return data
}

func Pub() {
	id := "sensor1"
	client := ConnectToMQTT(id)
	data := MessageSender(id)
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error converting data to JSON", err)
		return
	}
	msg := string(jsonData)
	PublishToBroker(client, "/sensor", msg)
	fmt.Println("Published:", msg)
	TimeToPublish()
}
