package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var num = "0"

func main() {
	e := echo.New()
	e.GET("/get_message", getMessage)
	e.POST("/verify", verifyWallet)
	e.Logger.Fatal(e.Start(":1323"))
}

func getMessage(c echo.Context) error {
	num1 := genRandomInt()
	num = strconv.Itoa(num1)

	return c.String(http.StatusOK, "message："+num)
}

func verifyWallet(c echo.Context) error {
	addr := c.FormValue("address")
	signedMsg := c.FormValue("signedMessage")

	res := Verify(num, addr, signedMsg)

	if res {
		return c.String(http.StatusOK, "verified: true")
	} else {
		return c.String(http.StatusOK, "verified: false")
	}
}
