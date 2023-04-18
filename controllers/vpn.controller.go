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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var vpnService *services.VpnService

func (ctl *VpnController) CrawlVpn(c echo.Context) error {
	param := c.QueryParam("url")

	// Kiểm tra xem chuỗi có phải là URL hợp lệ hay không
	_, err := url.ParseRequestURI(param)
	if err != nil {
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad Request!", "Chuỗi không phải là URL hợp lệ"))
	}
	bytesBody, err := helper.CallApi(param)

	// Eliminate redundant fields ==> csv format
	bytesBody = bytes.Trim(bytesBody, "*vpn_servers")
	bytesBody = bytesBody[:len(bytesBody)-1]
	s := strings.ReplaceAll(string(bytesBody), `"`, ``)
	reader := csv.NewReader(strings.NewReader(string(s)))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
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
		panic(err)
	}

	ResJsonData := make([]models.VpnModel, 0)
	err = json.Unmarshal(jsonData, &ResJsonData)
	if err != nil {
		panic(err)
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
			bytes, err := base64.StdEncoding.DecodeString(j.OpenVPN_ConfigData_Base64) // convert to base64
			err = ioutil.WriteFile("./config.ovpn", bytes, 0)                          // fill config vpn to file config.ovpn
			if err != nil {
				fmt.Println("\n error:", err)
				return c.HTML(http.StatusOK, "<b>Thank you! "+"</b>")

			}
			j.CreatedAt = time.Now()
			j.UpdatedAt = time.Now()
			if helper.CheckVpnIsLive() {
				j.Live = true
			} else {
				j.Live = false
			}
			vpnService.CreateVpn(j)
		}
	}
	return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", nil))

}

func (ctl *VpnController) GetAll(c echo.Context) error {
	live := c.QueryParam("live")
	query := false
	if live == "" || live == "false" {
		query = false
	} else if live == "true" {
		query = true
	}
	res, _ := vpnService.FindVpn(query)
	return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", res))

}

func (ctl *VpnController) Download(c echo.Context) error {
	name := c.Param("name")
	fmt.Print("----->>", name)
	res := vpnService.FindOneVpnByName(name)
	if res.OpenVPN_ConfigData_Base64 == "" {
		return c.JSON(http.StatusOK, utils.ResFail(404, "Not Found!", "Can't found vpn by "+name))

	}
	configbyte, err := base64.StdEncoding.DecodeString(res.OpenVPN_ConfigData_Base64)
	if err != nil {
		fmt.Println("Lỗi:", err.Error())
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad Request!", "can't download config"))
	}
	return c.Blob(http.StatusOK, "multipart/form-data", configbyte)

}
