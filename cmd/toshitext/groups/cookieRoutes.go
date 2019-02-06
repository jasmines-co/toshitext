package groups

import (
	"github.com/jasmines-co/toshitext/cmd/toshitext/handlers"
	"github.com/labstack/echo"
)

func CookieGroupRoutes(group *echo.Echo) {
	group.GET("/main", handlers.MainCookie)
}
