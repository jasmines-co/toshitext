package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type command struct {
	Help    string `json:"HELP"`
	Create  string `json:"CREATE"`
	Send    string `json:"SEND"`
	Receive string `json:"RECEIVE"`
	Buy     string `json:"BUY"`
	Sell    string `json:"SELL"`
}

func getHelpCommands(c echo.Context) error {
	cmd := command{
		Help:    "HELP",
		Create:  "CREATE",
		Send:    "SEND",
		Receive: "RECEIVE",
		Buy:     "BUY",
		Sell:    "SELL",
	}

	commands := []command{cmd}

	bs, err := json.Marshal(commands)
	if err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, string(bs))
	// fmt.Println(commands)
}
