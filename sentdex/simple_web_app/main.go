package main

import (
	"fmt"
	"net/http"
)

// IndexHandler handles index
func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintf(writer, "Whoa, Go is neat!")
}

//AboutHandler handles about page
func AboutHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Fprintf(writer, "We gonna be Go experts!!")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)

	http.ListenAndServe(":8000", nil)
}
