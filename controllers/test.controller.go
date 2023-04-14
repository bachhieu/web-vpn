package controllers

import (
	"bachhieu/web-vpn/services"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/labstack/echo/v4"
)

var testService = &services.TestService{}

func (ctl *TestController) Index(c echo.Context) error {
	cmd, err := exec.Command("ping", "8.8.8.8", "-w", "2").Output()
	if err != nil {
		fmt.Printf("err : %s", err)
	} else {
		fmt.Printf("result : %s", cmd)
	}
	string := testService.Index()
	return c.String(http.StatusOK, string)

}
