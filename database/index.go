package database

import (
	"bachhieu/web-vpn/services"
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var client *mongo.Client
// var vpnCollection *mongo.Collection

func Connection() {
	MONGO_URL := os.Getenv("MONGO_URL")
	fmt.Printf("\n MONGO_URL: %s \n", MONGO_URL)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		panic(err)
	}
	createCollection(client)
	fmt.Println("\nConnect to mongodb successful!")
}

// set up collection for each feature
func createCollection(client *mongo.Client) {
	services.CreateVpnCollectioon(client)
}
