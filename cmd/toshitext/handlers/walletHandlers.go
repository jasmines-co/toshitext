package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Wallet struct {
	Address string `json:address`
	Token   string `json:country`
}

func CreateWallet(c echo.Context) error {
	wallet := Wallet{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&wallet)
	if err != nil {
		log.Printf("Failed processing createWallet request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is your wallet address: %#v", wallet)
	return c.String(http.StatusOK, "We have successfully added a new wallet!")
}
