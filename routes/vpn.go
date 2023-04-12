package routes

import (
	"bachhieu/web-vpn/helper"
	"bachhieu/web-vpn/models"
	res "bachhieu/web-vpn/utils"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetAllVpns(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Call Api")

	resp, err := http.Get("http://www.vpngate.net/api/iphone/")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle error
		res.JSON(w, 400, err)


	}
	body = bytes.Trim(body, "*vpn_servers")
	body = body[:len(body)-1]
	// res.JSON(w, 400, body)
	// return
	
	s := strings.ReplaceAll(string(body), `"`, ``)
	reader := csv.NewReader(strings.NewReader(string(s)))
	fmt.Println("paser CSV")

	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		res.JSON(w, 400, err)

		return
	}


	var data []map[string]string
	for _, record := range records[1:] {
		row := make(map[string]string)
		for i, value := range record {
			row[records[0][i]] = value
		}
		data = append(data, row)
	}
	data = data[:len(data)-1]
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	
	ResJsonData := make([]models.Vpn, 0)
	err = json.Unmarshal(jsonData, &ResJsonData)
	if err != nil {
		panic(err)
	} 
	fmt.Println("loop ResJsonData ")
	for i, j := range ResJsonData {
		bytes, err := base64.StdEncoding.DecodeString(j.OpenVPN_ConfigData_Base64)
			err = ioutil.WriteFile("vpn.ovpn", bytes, 0)
			if err != nil {
				fmt.Println("error:", err)
				return
		}
		live := helper.CheckVpnIsLive()
	fmt.Printf("index: %v--- live: %v \n ",i,live)

	helper.KillProcess()
		if live == true {
			// save in DB
		} else {
			//edit file vpn.ovpn and rerun
			// save in DB
			continue
		}
	}
}