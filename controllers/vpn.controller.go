package controllers

import (
	"bachhieu/web-vpn/helper"
	"bachhieu/web-vpn/models"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (ctl *VpnController) CrawlVpn(c echo.Context) error {
	// start : Call api and format data
	resp, err := http.Get("http://www.vpngate.net/api/iphone/")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)

	}
	body = bytes.Trim(body, "*vpn_servers")
	body = body[:len(body)-1]
	s := strings.ReplaceAll(string(body), `"`, ``)
	reader := csv.NewReader(strings.NewReader(string(s)))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	//end
	// strart : convert type [][]string to  []models.Vpn
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
	// end
	for _, j := range ResJsonData {
		// here
		// check in databse if exist then next
		// if not exist check vpn is live
		//
		bytes, err := base64.StdEncoding.DecodeString(j.OpenVPN_ConfigData_Base64)
		err = ioutil.WriteFile("./config.ovpn", bytes, 0)
		if err != nil {
			fmt.Println("error:", err)
			return c.HTML(http.StatusOK, "<b>Thank you! "+"</b>")

		}
		if helper.CheckVpnIsLive() {
			// save in DB
		} else {
			//edit file vpn.ovpn and rerun
			// save in DB
			continue
		}
	}
	return c.HTML(http.StatusOK, "<b>Thank you! "+"</b>")

}
