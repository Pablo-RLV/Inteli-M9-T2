package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func PublishDatabase(data map[string]interface{}) {
	var client = ConnectClient()
	collection := client.Database("Ponderada").Collection("gas")
	document := bson.D{
		{"ID", data["ID"]},
		{"CO", data["CO"]},
		{"NO2", data["NO2"]},
		{"NH3", data["NH3"]},
	}
	insertResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)
	DisconnectClient(client)
}