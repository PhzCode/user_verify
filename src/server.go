package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/get_message", getMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func getMessage(c echo.Context) error {
	num1 := rand.Int()
	return c.String(http.StatusOK, strconv.Itoa(num1))
}