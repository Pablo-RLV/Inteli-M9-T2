package main

import (
	publisher "Prova1/publisher"
	subscriber "Prova1/subscriber"
)

func main() {
	for {
		publisher.Pub("l1", "freezer")
		subscriber.Sub()
		publisher.Pub("l2", "geladeira")
		subscriber.Sub()
	}
}
