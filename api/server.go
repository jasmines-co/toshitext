package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Wallet struct {
	Address string `json:address`
	Token   string `json:token`
}

func renderHome(c echo.Context) error {
	return c.String(http.StatusOK, "Server over here!")
}

func createWallet(c echo.Context) error {
	wallet := Wallet{}
	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&wallet)
	if err != nil {
		log.Printf("Failed processing createWallet request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("This is the wallet address: %#v", wallet)
	return c.String(http.StatusOK, "New crypto wallet created!")
}

////////////////////////// Middlewares ////////////////////////////

func main() {
	fmt.Println("Live from the server...")

	e := echo.New()

	// e.Use(ServerHeader)

	g := e.Group("/")

	// Log the server interaction
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.GET("/", renderHome)
	e.POST("/wallets", createWallet)

	e.Start(":8000")
}
