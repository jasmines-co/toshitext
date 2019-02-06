package router

import (
	handlers "github.com/jasmines-co/toshitext/cmd/toshitext/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/home", handlers.RenderHome)
	e.GET("/users/:id", handlers.GetUserById)
	e.GET("/wallet/:data", handlers.GetUserByWalletAddress)
	e.GET("/phones/:id", handlers.GetPhoneNumberForUser)
	e.GET("/phones", handlers.CheckIfPhoneNumberExists)
	e.GET("/exchange-rate", handlers.getExchangeRate)
	e.GET("/users/:id", handlers.GetBalance)
	e.GET("/users/:id", handlers.GetDepositAddress)
	e.GET("/users/:data", handlers.GetUser)
	e.GET("/wallets/:data", handlers.GetWalletAddress)

	e.POST("/users", handlers.AddUser)
	e.POST("/user:id", handlers.AddPhoneNumber)
	e.POST("/wallets", handlers.CreateWallet)
	e.POST("/user/:id", handlers.SellCrypto)
}
