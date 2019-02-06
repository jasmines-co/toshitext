package router

import (
	"projects/echo-server-template/api/handlers"

	"github.com/labstack/echo"
)

func CookkieGroup(g, *echo.Group) {
	g.GET("/main", handlers.GetMainCookie)
}
