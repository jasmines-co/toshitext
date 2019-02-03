package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CoinGecko struct {
	Rates struct {
		Name  string `json:name`
		Unit  string `json:unit`
		Value string `json:value`
		Type  string `json:value`
	}
}

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

func getBitcoinPrice(c echo.Context) error {
	rates := CoinGecko{}

	// Build the request
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/exchange_rates", nil)
	if err != nil {
		fmt.Println("Error requesting API endpoint: ", err)
	}

	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error getting bitcoin price: ", err)
	}

	// Defer the closing of the body
	defer response.Body.Close()

	// Use json.Decode for reading streams of JSON data
	// err := json.NewDecoder(response.Body).Decode(&employee)
	// if err != nil {
	// 	log.Printf("Failed processing createWallet request: %s", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError)
	// }

	// // Use json.Decode for reading streams of JSON data
	if error := json.NewDecoder(response.Body).Decode(&rates); error != nil {
		fmt.Println(error)
	}

	return c.JSON(http.StatusOK, rates)
}

////////////////////////// Middlewares ////////////////////////////

func main() {
	fmt.Println("Live from the server...")

	e := echo.New()

	// Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	g := e.Group("/")

	// Log the server interaction
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.GET("/", renderHome)
	e.POST("/wallets", createWallet)
	e.GET("/rates", getBitcoinPrice)

	e.Start(":8000")
}
