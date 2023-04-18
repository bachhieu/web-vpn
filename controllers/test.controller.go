package controllers

import (
	"bachhieu/web-vpn/services"
	"fmt"
	"net/http"
	"net/url"
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

// func (ctl *TestController) Query(c echo.Context) error {
// 	path := c.Param("path")
// 	paths := c.ParamNames()
// 	querys := c.QueryParams()
// 	query := c.QueryParam("path")
// 	fmt.Printf("paths------>%s \n", paths)
// 	fmt.Printf("querys------>%s \n", querys)
// 	fmt.Printf("query------>%s \n", query)
// 	return c.String(http.StatusOK, path)

// }

func (ctl *TestController) Query(c echo.Context) error {
	param := c.QueryParam("url")

	// Kiểm tra xem chuỗi có phải là URL hợp lệ hay không
	_, err := url.ParseRequestURI(param)
	if err != nil {
		fmt.Println("Chuỗi không phải là URL hợp lệ")
	} else {
		fmt.Println("Chuỗi là URL hợp lệ")
	}
	return c.String(http.StatusOK, param)

}
