package main

import (
	publisher "Ponderada2/publisher"
	subscriber "Ponderada2/subscriber"
)

func main() {
	for{
		publisher.Pub()
		subscriber.Sub()
	}
}
