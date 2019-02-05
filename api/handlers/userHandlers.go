package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
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
