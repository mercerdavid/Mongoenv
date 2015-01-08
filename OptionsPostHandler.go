// GetHandler
package main

import (
	"fmt"
	"net/http"
)

func OptionsPostHandler(res http.ResponseWriter, req *http.Request) {
	//var role bubbleRole
	fmt.Println("Origin OptionsPostHandler >>" + req.Header.Get("Origin"))
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
	}

	fmt.Println("in post handler ")
	res.Header().Set("Content-Type", "text/json")

	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	res.Header().Set("Access-Control-Allow-Credentials", "true")

	fmt.Println("past headers")

	res.Write([]byte("hello2"))
}
