package routes

import (
	"bachhieu/web-vpn/controllers"

	"github.com/labstack/echo/v4"
)

var testController = &controllers.TestController{}

func TestInit(g *echo.Group) {
	g.GET("/", testController.Index)
	g.GET("/query", testController.Query)
}
