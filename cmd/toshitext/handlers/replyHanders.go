package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	plivo "github.com/plivo/plivo-go"
)

type Reply struct {
	ID          int    `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
}

func ReplyToMessage(c echo.Context) error {
	plivoAuthId := os.Getenv("PLIVO_AUTH_ID")
	plivoAuthToken := os.Getenv("PLIVO_AUTH_TOKEN")
	client, err := plivo.NewClient(plivoAuthId, plivoAuthToken, &plivo.ClientOptions{})
	if err != nil {
		panic(err)
	}

	response, err := client.Messages.Create(plivo.MessageCreateParams{
		Src:  "+14155696002",
		Dst:  "+17025305234",
		Text: "Toshitext just sent some crypto 😎!",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Response: %#v\n", response)
	return c.String(http.StatusOK, "We have successfully sent the message!")
}
