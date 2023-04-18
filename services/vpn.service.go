package services

import (
	"bachhieu/web-vpn/models"
	"bachhieu/web-vpn/utils"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var vpnCollection *mongo.Collection

func CreateVpnCollectioon(client *mongo.Client) {
	DATABASE := os.Getenv("MONGO_INITDB_DATABASE")
	fmt.Printf("\n DATABASE: %s \n", DATABASE)
	vpnCollection = client.Database(DATABASE).Collection("vpn")
}

func (ctl *VpnService) FindVpn(query bool) (utils.ResData, error) {
	cur, err := vpnCollection.Find(context.Background(), bson.M{"live": query})
	defer cur.Close(context.Background())
	vpnModel := []models.VpnModel{}
	err = cur.All(context.Background(), &vpnModel)
	if err != nil {
		return utils.ResData{Total: 0, Data: nil}, err
	}
	return utils.ResData{Total: len(vpnModel), Data: vpnModel}, err
}

func (ctl *VpnService) CheckVpnExistAndLive() models.VpnModel {
	cur, err := vpnCollection.Find(context.Background(), bson.M{"live": true})
	defer cur.Close(context.Background())
	vpnModel := models.VpnModel{}
	err = cur.Decode(&vpnModel)
	if err != nil {
		log.Fatal(err)
	}
	return vpnModel
}

func (ctl *VpnService) CreateVpn(vpn models.VpnModel) bool {
	_, err := vpnCollection.InsertOne(context.Background(), vpn)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (ctl *VpnService) CheckVpnIsExistByName(query string) bool {
	var vpn models.VpnModel
	cur := vpnCollection.FindOne(context.Background(), bson.M{"hostname": query})
	err := cur.Decode(&vpn)

	if err != nil {
		return false
	}
	return true
}

func (ctl *VpnService) FindOneVpnByName(query string) models.VpnModel {
	var vpn models.VpnModel
	cur := vpnCollection.FindOne(context.Background(), bson.M{"hostname": query})
	err := cur.Decode(&vpn)

	if err != nil {
		return models.VpnModel{}
	}
	return vpn
}
