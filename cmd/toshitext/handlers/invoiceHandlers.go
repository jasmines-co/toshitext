package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

type Invoice struct {
	ID            int `json:id`
	Currency      int `json:currency`
	PaymentAmount int `json:paymentAmount`
}

func GetInvoice(c echo.Context) error {
	return c.String(http.StatusOK, "Display invoice info")
}
