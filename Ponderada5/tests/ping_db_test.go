package tests

import (
	"testing"
	database "Ponderada5/database"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func TestPing(t *testing.T) {
	var client = database.ConnectClient()
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	database.DisconnectClient(client)
}
