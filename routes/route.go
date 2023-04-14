package routes

import (
	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	VpnInit(g.Group("/vpn"))   // set route api/v1/vpn
	TestInit(g.Group("/test")) // set route api/v1/test for test
}
