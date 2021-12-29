package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAccess(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func PostAccess(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
