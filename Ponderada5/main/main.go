package main

import (
	publisher "Ponderada5/publisher"
	subscriber "Ponderada5/subscriber"
)

func main() {
	go publisher.Pub("sensor1")
	subscriber.Sub("subscriber")
	select {}
}