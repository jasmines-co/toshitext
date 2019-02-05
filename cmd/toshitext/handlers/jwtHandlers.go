package handlers

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetMainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("Username: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "You are on the correct JWT page!")
}
