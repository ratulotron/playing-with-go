package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// Lowercase properties are not exported in Go.
	// As a result we have to capitalize the first letter of the prop
	// So if we want lowercased proper json response we have to use
	// field attributes.
	Message string `json:"message"`
	Author  string `json:"-"`          // no output
	Date    string `json:",omitempty"` // no output if empty
	ID      int    `json:"id,string"`  // convert this
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{
		Message: "Hello World",
		Author:  "Ratul",
		ID:      20,
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
}

func main() {
	port := 8080
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
