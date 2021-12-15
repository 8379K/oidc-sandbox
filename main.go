package main

import (
	"fmt"

	"github.com/8379K/oidc-sandbox/handler"
	"github.com/labstack/echo/v4"
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

		}
	}
}
