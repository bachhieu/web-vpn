package routes

import (
	"bachhieu/web-vpn/controllers"

	"github.com/labstack/echo/v4"
)

var vpnController = &controllers.VpnController{}

func VpnInit(g *echo.Group) {
	g.GET("/vpngate/crawl", vpnController.CrawlVpngate)
	g.GET("/vpngate/crawl/toggle", vpnController.CrawlVpngate)
	g.GET("/get-all", vpnController.GetAll)
	g.GET("/:name/download", vpnController.Download)
}
