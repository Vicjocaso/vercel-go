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
	"fmt"
	"io"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go RunTime Serverless!</h1>")
}

func GetData(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>DATA!</h1>")
}

func main() {

	// Start an HTTP API server.
	const addr = ":3000"
	log.Printf("successfully connected to PlanetScale, starting HTTP server on %q", addr)
	http.HandleFunc("/", Handler)
	http.HandleFunc("/get-data", GetData)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	io.WriteString(w, `{"alive": true}`)
}
