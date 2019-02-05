package api

import (
	"projects/echo-server-template/api/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/home", handlers.RenderHome)
	e.GET("/login", handlers.Login)
	e.GET("/users/:data", handlers.GetUser)
	e.GET("/wallets/:data", handlers.GetWalletAddress)

	e.POST("/users", handlers.AddUser)
	e.POST("/phones", handlers.AddPhone)
	e.POST("/wallets", handlers.CreateWallet)
}
