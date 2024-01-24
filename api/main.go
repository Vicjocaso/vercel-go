package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func handleEndpoint1(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 1!")
}

func handleEndpoint2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from endpoint 2!")
}

func main() {
	e := echo.New()

	e.GET("/endpoint1", handleEndpoint1)
	e.GET("/endpoint2", handleEndpoint2)

	e.Logger.Fatal(e.Start(":3000"))
}
