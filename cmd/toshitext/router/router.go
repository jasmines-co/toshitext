package router

import (
	"github.com/jasmines-co/toshitext/cmd/toshitext/groups"
	"github.com/jasmines-co/toshitext/cmd/toshitext/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// Set all middlewares
	middlewares.SetMainMiddlewares(e)

	// Set main routes
	groups.MainGroup(e)

	return e
}
