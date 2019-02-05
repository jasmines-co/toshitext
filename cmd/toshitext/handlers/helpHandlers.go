package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Help struct {
	ID int
}

func GetMainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "Text HELP to get list of commands")
}
