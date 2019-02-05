package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetMainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "You are on the secret cookie page.")
}
