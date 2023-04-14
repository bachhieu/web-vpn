package models

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Connection() {
	MONGO_URL := os.Getenv("MONGO_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		panic(err)
	}
}

func GetCollection(collection string) *mongo.Collection {
	return client.Database("vpn").Collection(collection)
}
