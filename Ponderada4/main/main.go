package main

import (
	publisher "Ponderada4/publisher"
	subscriber "Ponderada4/subscriber"
)

func main() {
	for {
		publisher.Pub()
		subscriber.Sub()
	}
}
