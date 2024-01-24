package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 1!")
}

func HandleEndpoint2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 2!")
}

func main() {
	e := echo.New()

	e.GET("/1", Handler)
	e.GET("/2", HandleEndpoint2)

	e.Logger.Fatal(e.Start(":3000"))
}
