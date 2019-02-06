package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type PhoneNumber struct {
	ID      int
	User    *User
	Phone   string
	Country string
	Status  string
}

func GetPhoneNumbers(c echo.Context) error {
	return c.String(http.StatusOK, "Display Phone Info")
}
