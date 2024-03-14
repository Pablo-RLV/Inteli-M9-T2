package database

import (
	common "Ponderada5/common"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func connect_client() *mongo.Client {
	common.LoadEnv()
	var connectionString = os.Getenv("CONNECTION_STRING")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return client
}

func disconnect_client(client *mongo.Client) {
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func PingDatabase(client *mongo.Client) {
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func PublishDatabase(data map[string]interface{}) {
	var client = connect_client()
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
	disconnect_client(client)
}