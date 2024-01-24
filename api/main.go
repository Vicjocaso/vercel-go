package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleEndpoint1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 1!")
}

func HandleEndpoint2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 2!")
}

func main() {
	e := echo.New()

	e.GET("/endpoint1", HandleEndpoint1)
	e.GET("/endpoint2", HandleEndpoint2)

	e.Logger.Fatal(e.Start(":3000"))
}
