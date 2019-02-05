package handlers

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:name`
	jwt.StandardClaims
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"fode",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("Not_So_Secret"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	// check username nd password in DB after hash
	if username == "foo" && password == "bar" {
		cookie := &http.Cookie{}

		cookie.Name = "sessionID"
		cookie.Value = "random_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(cookie)

		// TODO: create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Error creating JWT token", err)
			return c.String(http.StatusInternalServerError, "Something went  wrong")
		}

		jwtCookie := &http.Cookie{}

		jwtCookie.Name = "JWTCookie"
		jwtCookie.Value = token
		jwtCookie.Expires = time.Now().Add(48 * time.Hour)

		c.SetCookie(jwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were succesfuly logged in",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Wrong username or password")
}
