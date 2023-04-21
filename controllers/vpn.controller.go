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
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

var vpnService *services.VpnService
var cronJob *cron.Cron
var _ = godotenv.Load()
var env = struct {
	VPNGATE string
	CRONJOB string
}{
	os.Getenv("VPNGATE"),
	os.Getenv("CRONJOB"),
}
var idCrawlVpn cron.EntryID

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
		idCrawlVpn = utils.AddFunc(cronJob, "0 12 * * *", crawlVpngateHelper)
		return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Switch crawl vpn auto successful!"))
	case !auto && VPNGATE:
		utils.RemoveFunc(cronJob, idCrawlVpn)
		c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Switch crawl vpn manual successful!"))
		defer crawlVpngateHelper()
		return nil
	}
	defer crawlVpngateHelper()
	return nil

}

func (ctl *VpnController) GetVpnByLive(c echo.Context) error {
	live := c.QueryParam("live")
	query := false
	if live == "" || live == "false" {
		query = false
	} else if live == "true" {
		query = true
	}
	m := map[string]interface{}{
		"live": query,
	}
	result := vpnService.FindVpn(m)

	res := utils.ResData{Total: len(result), Data: result}
	return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", res))

}

func (ctl *VpnController) Download(c echo.Context) error {
	ctl.cronVpnHelper()
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

func (ctl *VpnController) ToggleCronjob(c echo.Context) error {
	cron := c.QueryParam("cron") == "true"
	switch {
	case cron && env.CRONJOB == "true":
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad request!", "Schedule really is runing!"))
	case !cron && env.CRONJOB == "true":
		cronJob.Stop()
		env.CRONJOB = "false"
		return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Stop cron job successful!"))
	case cron && env.CRONJOB == "false":
		cronJob.Start()
		env.CRONJOB = "true"
		return c.JSON(http.StatusOK, utils.ResSuccess(200, "Successfull!", "Run cron job successful!"))
	default:
		return c.JSON(http.StatusOK, utils.ResFail(400, "Bad request!", "Schedule really is off!"))
	}
}

func (ctl *VpnController) CronVpn(schedule string) {
	cronJob = utils.Schedule(schedule, ctl.cronVpnHelper)
	if os.Getenv("VPNGATE") == "true" {
		idCrawlVpn = utils.AddFunc(cronJob, "@daily", crawlVpngateHelper)
	}
}

func (ctl *VpnController) cronVpnHelper() {
	m := map[string]interface{}{}
	vpns := vpnService.FindVpn(m)
	size, err := strconv.Atoi(os.Getenv("NUM_GOROUTINES"))
	if err != nil {
		size = 5 // default 5 goroutines
	}
	numJobs := size
	len := len(vpns)
	jobs := make(chan models.VpnModel, len)
	results := make(chan models.VpnModel, len)
	for w := 1; w <= numJobs; w++ {
		go scanCheckVpnIsLive(jobs, results)
	}
	for _, j := range vpns {
		jobs <- j
	}
	close(jobs)

	for _, i := range vpns {
		vpn := <-results

		for _, k := range vpns {
			if k != i {

			}
		}
		vpnService.UpdatedOne(vpn)
	}

}

func crawlVpngateHelper() {
	bytesBody, err := helper.CallApi("https://www.vpngate.net/api/iphone/")

	// Eliminate redundant fields ==> csv format
	bytesBody = bytes.Trim(bytesBody, "*vpn_servers")
	bytesBody = bytesBody[:len(bytesBody)-1]
	s := strings.ReplaceAll(string(bytesBody), `"`, ``)
	reader := csv.NewReader(strings.NewReader(string(s)))
	reader.FieldsPerRecord = -1
	records, err := reader.ReadAll()
	if err != nil {
		return
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
		return
	}

	ResJsonData := make([]models.VpnModel, 0)
	err = json.Unmarshal(jsonData, &ResJsonData)
	if err != nil {
		return
	}
	// end
	for i, j := range ResJsonData {
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
	return

}

func scanCheckVpnIsLive(vpns <-chan models.VpnModel, result chan<- models.VpnModel) {
	for i := range vpns {
		bytes, _ := base64.StdEncoding.DecodeString(i.OpenVPN_ConfigData_Base64) // convert to base64
		if helper.CheckVpnIsLive(bytes) {
			i.UpdatedAt = time.Now()
			result <- i
		} else {
			if i.Live == false {
				result <- i
			} else {
				i.UpdatedAt = time.Now()
				result <- i
			}
		}
	}
}

func setupConcu(vpns <-chan models.VpnModel, result chan<- models.VpnModel) {
	for i := range vpns {
		bytes, _ := base64.StdEncoding.DecodeString(i.OpenVPN_ConfigData_Base64) // convert to base64
		if helper.CheckVpnIsLive(bytes) {
			i.UpdatedAt = time.Now()
			result <- i
		} else {
			if i.Live == false {
				result <- i
			} else {
				i.UpdatedAt = time.Now()
				result <- i
			}
		}
	}
}
