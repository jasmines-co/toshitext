package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Reply struct {
	ID          int    `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
}

func replyToMessage() {
	// Initialize godotenv for reading secrets stored in .env files.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
}
