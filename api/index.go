// package handler

// import (
// 	"fmt"
// 	"net/http"
// )

// type Film struct {
// 	Title    string
// 	Director string
// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "<h1>Hello from Go RunTime Serverless!</h1>")
// }

// func GetData(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprintf(w, "<h1>Hello from Go RunTime Serverless!</h1>")
// }

// func main() {
// 	http.HandleFunc("/", Handler)
// 	http.HandleFunc("/get-data", GetData)
// 	http.ListenAndServe(":3000", nil)
// }

package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()

	e.GET("/endpoint1", func(c echo.Context) error {
		return c.String(http.StatusOK, "<h1>Hello from Go RunTime Serverless!</h1>")
	})

	e.GET("/endpoint2", func(c echo.Context) error {
		return c.String(http.StatusOK, "<h1>Hello from Go RunTime Serverless 2!</h1>")
	})

	e.Logger.Fatal(e.Start(":3000"))

}
