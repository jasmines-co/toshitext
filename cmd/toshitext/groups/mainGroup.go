package groups

import (
	handlers "github.com/jasmines-co/toshitext/cmd/toshitext/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/home", handlers.RenderHome)
	e.POST("/wallets", handlers.CreateWallet)
}
