// mongoserv project main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	// set content type to html example
	r.HandleFunc("/get", GetHandler).Methods("GET")
	r.HandleFunc("/post", PostHandler).Methods("POST")
	r.HandleFunc("/post", OptionsPostHandler).Methods("OPTIONS")

	//-
	http.Handle("/", r)

	fmt.Println("Listening on port 8083")

	// wait for clients
	http.ListenAndServe(":8083", nil)
}
