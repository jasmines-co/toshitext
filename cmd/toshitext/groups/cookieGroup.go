package groups

import (
	"github.com/jasmines-co/toshitext/cmd/toshitext/handlers"

	"github.com/labstack/echo"
)

func CookkieGroup(g, *echo.Group) {
	g.GET("/main", handlers.GetMainCookie)
}
