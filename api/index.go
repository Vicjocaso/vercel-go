package handler

import (
	"net/http"
	"vercel-golang/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handlers.GetHandleFunc(w, r)
	case "POST":
		handlers.PostHandleFunc(w, r)
	case "PUT":
		handlers.PutHandleFunc(w, r)
	default:
		return
	}
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":3000", nil)
}
