package handler

import (
	"fmt"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go RunTime Serverless!</h1>")
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}
