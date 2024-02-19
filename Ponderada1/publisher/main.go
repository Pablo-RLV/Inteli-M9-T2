package main

import (
	"encoding/json"
	"fmt"
	"time"
	"math/rand"

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
			"CO_ppm": rand.Intn(1000),
			"NO2_ppm": rand.Intn(10),
			"NH3_ppm": rand.Intn(500),
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
