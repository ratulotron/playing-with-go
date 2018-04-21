package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	port := 8080

	// Creates a Handler type on the DefaultServeMux handler
	// mapping the path passed
	// to the function
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	// Start the HTTP server ListenAndServe
	// One param is the port number,
	// Second one is nil coz we are using default serve mux handler
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
