package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// func HandleEndpoint2(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello from endpoint 2!")
// }

// func main() {
// 	e := echo.New()

// 	e.GET("/endpoint1", Handler)
// 	e.GET("/endpoint2", HandleEndpoint2)

// 	e.Logger.Fatal(e.Start(":3000"))
// }
