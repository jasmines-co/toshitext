package main

import (
	"fmt"
	"router"
)

func main() {
	fmt.Println("Live from the server...")

	e := router.New()

	e.Start(":8000")
}
