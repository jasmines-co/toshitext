package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Phone struct {
	Number  string `json:number`
	Country string `json:country`
}

func AddPhone(c echo.Context) error {
	phone := Phone{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&phone)
	if err != nil {
		log.Printf("Failed processing addPhoneNumber request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is the user: %#v", phone)
	return c.String(http.StatusOK, "We have added a new phone number!")
}
