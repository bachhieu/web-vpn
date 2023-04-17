package routes

import (
	"bachhieu/web-vpn/controllers"

	"github.com/labstack/echo/v4"
)

var vpnController = &controllers.VpnController{}

func VpnInit(g *echo.Group) {
	g.GET("/crawl", vpnController.CrawlVpn)
	g.GET("/get-all", vpnController.GetAll)
}
