package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	ID            int    `json:id`
	Username      string `json:username`
	FirstName     string `json:firstName`
	LastName      string `json:firstName`
	Status        string `json:status`
	Balance       string `json:balance`
	WalletAddress string `json:walletAddress`
}

func AddUser(c echo.Context) error {
	user := User{}

	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body for addUser: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &user)
	if err != nil {
		log.Printf("Failed unmarshaling in addUser : %s", err)
	}

	log.Printf("This is the user: %#v", user)
	return c.String(http.StatusOK, "We have a user!")

}
