package main

import (
	publisher "Ponderada7/publisher"
	kafka "Ponderada7/kafka"
)

func main() {
	go publisher.Pub("sensor3")
	kafka.Kafka()
	select {}
}
