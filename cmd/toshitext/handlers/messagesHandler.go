package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
	ID          int    `json:id`
	Recipient   string `json:id`
	MessageTime string `json:messageTime`
	Status      string `json:status`
	Text        string `json:text`
	FirstName   string `json:firstName`
	LastName    string `json:lastName`
	Country     string `json:country`
	Sender      string `json:sender`
	Spam        bool   `json:spam`
}

func GetMessages(c echo.Context) error {
	return c.String(http.StatusOK, "Messages live here.")
}
