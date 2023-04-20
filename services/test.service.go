package services

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func (ctl *TestService) Index() string {
	fmt.Println("<b>Thank you! " + "example" + "</b>")
	return "<b>Thank you! " + "example" + "</b>"

}

var testCollection *mongo.Collection

func CreateTestCollectioon(client *mongo.Client) {
	DATABASE := os.Getenv("MONGO_INITDB_DATABASE")
	fmt.Printf("\n DATABASE: %s \n", DATABASE)
	testCollection = client.Database(DATABASE).Collection("test")
	test()
}

func test() {
	vpnCollection.InsertOne(context.Background(), struct {
		Test string `json:"test"`
	}{
		"test",
	})
}
