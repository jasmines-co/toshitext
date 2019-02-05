package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Contact struct {
	ID        int    `json:id`
	Phone     string `json:phone`
	FirstName string `json:firstName`
	LastName  string `json:lastName`
	Country   string `json:country`
}

func GetContacts(c echo.Context) error {
	return c.String(http.StatusOK, "Display Contacts Information for User")
}
