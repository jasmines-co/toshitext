package api

import "github.com/labstack/echo"

func CookieGroupRoutes(group *echo.Echo) {
	group.GET("/main", handlers.MainCookie)
}
