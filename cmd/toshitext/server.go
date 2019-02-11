package main

import (
	"fmt"

	"github.com/jasmines-co/toshitext/cmd/toshitext/router"
)

func main() {
	fmt.Println("Toshitext vAlpha - Send crypto with a text message.")

	e := router.New()

	e.Start(":8000")
}
