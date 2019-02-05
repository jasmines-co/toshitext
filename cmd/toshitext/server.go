package main

import (
	"fmt"

	"github.com/jasmines-co/toshitext/cmd/toshitext/router"
)

func main() {
	fmt.Println("Live from the server...")

	e := router.New()

	e.Start(":8000")
}
