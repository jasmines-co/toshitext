package groups

import handlers "github.com/jasmines-co/toshitext/cmd/toshitext/handlers"

func JwtGroup(g, *echo.Grop) {
	g.GET("/main", handlers.GetMainJwt)
}
