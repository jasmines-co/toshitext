package api

import (
	"projects/echo-server-template/api/handlers"

	"github.com/labstack/echo"
)

func AdminGroup(g, *echo.Group) {
	g.GET("/main", handlers.GetMainAdmin)
}
