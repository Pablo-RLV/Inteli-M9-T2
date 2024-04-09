package tests

import (
	database "Ponderada7/database"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestPing(t *testing.T) {
	var client = database.ConnectClient()
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	database.DisconnectClient(client)
}
