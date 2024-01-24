// package handler

// import (
// 	"fmt"
// 	"net/http"
// )

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
// }

//	func main() {
//		http.HandleFunc("/", Handler)
//		http.ListenAndServe(":3000", nil)
//	}
package handler

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
