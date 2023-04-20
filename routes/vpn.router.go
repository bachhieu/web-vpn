package routes

import (
	"bachhieu/web-vpn/controllers"

	"github.com/labstack/echo/v4"
)

var vpnController = &controllers.VpnController{}

func VpnInit(g *echo.Group) {
	g.GET("/vpngate/crawl", vpnController.CrawlVpngate)
	g.POST("/vpngate/crawl/toggle", vpnController.CrawlVpngate)
	g.POST("/cron/toggle", vpnController.ToggleCronjob)
	g.GET("/get-all", vpnController.GetVpnByLive)
	g.GET("/:name/download", vpnController.Download)
}
