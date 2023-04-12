package models

type Vpn struct {
	HostName                  string `json:"#HostName"`
	CountryLong               string `json:"CountryLong"`
	CountryShort              string `json:"CountryShort"`
	IP                        string `json:"IP"`
	LogType                   string `json:"LogType"`
	NumVpnSessions            string `json:"NumVpnSessions"`
	Operator                  string `json:"Operator"`
	Ping                      string `json:"Ping"`
	Score                     string `json:"Score"`
	Speed                     string `json:"Speed"`
	TotalTraffic              string `json:"TotalTraffic"`
	TotalUsers                string `json:"TotalUsers"`
	Uptime                    string `json:"Uptime"`
	Message                   string `json:"Message"`
	OpenVPN_ConfigData_Base64 string `json:"OpenVPN_ConfigData_Base64"`
}
