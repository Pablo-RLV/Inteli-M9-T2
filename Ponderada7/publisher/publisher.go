package publisher

import (
	common "Ponderada7/common"
	database "Ponderada7/database"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func PublishToBroker(client MQTT.Client, topic string, msg string) {
	token := client.Publish(topic, 1, false, msg)
	token.Wait()
}

func TimeToPublish() {
	time.Sleep(2 * time.Second)
}

func MessageFormat(clientID string) map[string]interface{} {
	data := map[string]interface{}{
		"ID":  clientID,
		"CO":  strconv.Itoa(rand.Intn(1000)) + " ppm",
		"NO2": strconv.Itoa(rand.Intn(10)) + " ppm",
		"NH3": strconv.Itoa(rand.Intn(500)) + " ppm",
	}
	return data
}

func Pub(id string) {
	client := common.ConnectToMQTT(id, nil)
	for {
		data := MessageFormat(id)
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}
		msg := string(jsonData)
		PublishToBroker(client, "/sensor", msg)
		database.PublishDatabase(data)
		fmt.Println("Published:", msg)
		TimeToPublish()
	}
}
