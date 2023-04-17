package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VpnModel struct {
	Id                        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt                 time.Time          `bson:"created_at,omitempty"`
	UpdatedAt                 time.Time          `bson:"updated_at,omitempty"`
	Live                      bool               `json:"live"`
	HostName                  string             `json:"#HostName"`
	CountryLong               string             `json:"CountryLong"`
	CountryShort              string             `json:"CountryShort"`
	IP                        string             `json:"IP"`
	LogType                   string             `json:"LogType"`
	NumVpnSessions            string             `json:"NumVpnSessions"`
	Operator                  string             `json:"Operator"`
	Ping                      string             `json:"Ping"`
	Score                     string             `json:"Score"`
	Speed                     string             `json:"Speed"`
	TotalTraffic              string             `json:"TotalTraffic"`
	TotalUsers                string             `json:"TotalUsers"`
	Uptime                    string             `json:"Uptime"`
	Message                   string             `json:"Message"`
	OpenVPN_ConfigData_Base64 string             `json:"OpenVPN_ConfigData_Base64"`
}
