package utils

import "github.com/labstack/echo/v4"

var skipPaths = []string{
	"api/docs/*",
	"/api/docs/*",
	"/api/docs",
}
var Skipper = func(c echo.Context) bool {
	for _, path := range skipPaths {
		if c.Path() == path {
			return true
		}
	}
	return false
}
