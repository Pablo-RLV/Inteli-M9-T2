package publisher

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"strconv"
	"time"
	godotenv "github.com/joho/godotenv"
	"os"
)

func ConnectToMQTT(clientID string) MQTT.Client {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883

	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
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
