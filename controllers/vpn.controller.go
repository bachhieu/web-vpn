package controllers

import (
	"bachhieu/web-vpn/helper"
	"bachhieu/web-vpn/models"
	"bachhieu/web-vpn/services"
	"bachhieu/web-vpn/utils"
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var vpnService *services.VpnService

func (ctl *VpnController) CrawlVpngate(c echo.Context) error {
	paramAuto := c.QueryParam("auto")
	if paramAuto != "true" && paramAuto != "false" {
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad Request!", "Field auto wrong format!"))
	}

	auto := paramAuto == "true"
	VPNGATE := env.VPNGATE == "true"
	utils.EditFile(".env", []byte("VPNGATE="+paramAuto), "VPNGATE=")
	env.VPNGATE = paramAuto
	switch {
	case auto && VPNGATE:
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad Request!", "Crawl data really is auto!"))
	case auto && !VPNGATE:
		return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Switch crawl vpn auto successful!"))
	case !auto && VPNGATE:
		defer c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Switch crawl vpn manual successful!"))
		return crawlVpngate(c)
	}
	return crawlVpngate(c)

}

func (ctl *VpnController) GetAll(c echo.Context) error {
	live := c.QueryParam("live")
	query := false
	if live == "" || live == "false" {
		query = false
	} else if live == "true" {
		query = true
	}
	result := vpnService.FindVpn(query)

	res := utils.ResData{Total: len(result), Data: result}
	return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", res))

}

func (ctl *VpnController) Download(c echo.Context) error {
	name := c.Param("name")
	res := vpnService.FindOneVpnByName(name)
	if res.OpenVPN_ConfigData_Base64 == "" {
		return c.JSON(http.StatusOK, utils.ResFail(404, "Not Found!", "Can't found vpn by "+name))

	}
	configbyte, err := base64.StdEncoding.DecodeString(res.OpenVPN_ConfigData_Base64)
	if err != nil {
		fmt.Println("Lỗi:", err.Error())
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad Request!", "can't download config of "+name))
	}
	return c.Blob(http.StatusOK, "multipart/form-data", configbyte)

}

func crawlVpngate(c echo.Context) error {
	bytesBody, err := helper.CallApi("https://www.vpngate.net/api/iphone/")

	// Eliminate redundant fields ==> csv format
	bytesBody = bytes.Trim(bytesBody, "*vpn_servers")
	bytesBody = bytesBody[:len(bytesBody)-1]
	s := strings.ReplaceAll(string(bytesBody), `"`, ``)
	reader := csv.NewReader(strings.NewReader(string(s)))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
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
		return err
	}

	ResJsonData := make([]models.VpnModel, 0)
	err = json.Unmarshal(jsonData, &ResJsonData)
	if err != nil {
		return err
	}
	// end
	for i, j := range ResJsonData {
		// here
		// check in databse if exist then next
		// if not exist check vpn is live
		fmt.Print(i, "-----")
		if vpnService.CheckVpnIsExistByName(j.HostName) {
			fmt.Print("------> next \n")
			continue
		} else {
			fmt.Print("------> import \n")
			bytes, err := base64.StdEncoding.DecodeString(j.OpenVPN_ConfigData_Base64) // convert to base64
			if err != nil {
				continue
			}
			j.CreatedAt = time.Now()
			j.UpdatedAt = time.Now()
			if helper.CheckVpnIsLive(bytes) {
				j.Live = true
			} else {
				j.Live = false
			}
			vpnService.CreateVpn(j)
		}
	}
	return nil

}
