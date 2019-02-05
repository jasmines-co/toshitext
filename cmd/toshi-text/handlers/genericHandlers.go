package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func RenderHome(c echo.Context) error {
	return c.String(http.StatusOK, "Server over here!")
}
