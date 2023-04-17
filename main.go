package main

import (
	"bachhieu/web-vpn/database"
	"bachhieu/web-vpn/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("\n Error loading .env file-->", err)
	}
	database.Connection() // connect to mongoDb
	// utils.Init() // enable tuntap run in docker
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.Init(e.Group("/api/v1")) // config routes for product

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
