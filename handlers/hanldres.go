package handlers

import (
	"fmt"
	"net/http"
)

func GetHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello from Go GET!</h1>")

}

func PostHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello from Go POST!</h1>")

}

func PutHandleFunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello from Go PUT!</h1>")

}
