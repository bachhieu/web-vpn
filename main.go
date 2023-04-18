package main

import (
	"bachhieu/web-vpn/database"
	_ "bachhieu/web-vpn/docs"
	"bachhieu/web-vpn/routes"
	"bachhieu/web-vpn/utils"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("\n Error loading .env file-->", err)
	}
	database.Connection() // connect to mongoDb
	// utils.Init() // enable tuntap run in docker
	e := echo.New()

	e.GET("/api/docs/*", echoSwagger.WrapHandler)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:  "method=${method}, uri=${uri}, status=${status}\n",
		Skipper: utils.Skipper,
	}))
	e.Use(middleware.Recover())
	routes.Init(e.Group("/api/v1")) // config routes for product

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
