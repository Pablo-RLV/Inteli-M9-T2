package main

import (
	"encoding/json"
	"fmt"
	"time"
	"math/rand"
	"strconv"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	id := "sensor1"
	opts := MQTT.NewClientOptions().AddBroker("tcp://broker:1891")
	opts.SetClientID(id)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} 

	for {
		data := map[string]interface{}{
			"ID": id,
			"CO": strconv.Itoa(rand.Intn(1000)) + " ppm",
			"NO2": strconv.Itoa(rand.Intn(10)) + " ppm",
			"NH3": strconv.Itoa(rand.Intn(500)) + " ppm",
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData)

		token := client.Publish("/sensor", 0, false, msg)
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}
