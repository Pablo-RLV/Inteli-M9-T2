package common

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"fmt"
	"os"
)

func SetOptions(clientID string) *MQTT.ClientOptions {
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID(clientID)
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	return opts
}