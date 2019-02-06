package groups

import (
	handlers "github.com/jasmines-co/toshitext/cmd/toshitext/handlers"
	"github.com/labstack/echo"
)

func AdminGroup(g, *echo.Group) {
	g.GET("/main", handlers.GetMainAdmin)
}
