package main

import (
	"fmt"
	"os"

	"github.com/8379K/oidc-sandbox/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("")
	e := echo.New()
	e.Static("/", "dist")

	api := e.Group("/api")
	{
		apiAccess := api.Group("/acess")
		{
			apiAccess.GET("", handler.GetAccess)
			apiAccess.POST("", handler.PostAccess)
		}
	}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
