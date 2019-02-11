package main

import (
	"fmt"
	"log"

	"github.com/jasmines-co/toshitext/cmd/toshitext/router"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize godotenv for reading secrets stored in .env files.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Toshitext vAlpha - Send crypto with a text message.")

	e := router.New()

	e.Start(":8000")
}
