package api

func JwtGroup(g, *echo.Grop) {
	g.GET("/main", handlers.GetMainJwt)
}
