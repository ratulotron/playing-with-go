package main

import (
	"net/http"

	"encoding/json"

	"log"

	"github.com/gorilla/mux"
)

// Person is a human
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address is a person's address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people := append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}

func main() {
	people = append(people, Person{
		ID:        "1",
		Firstname: "Ratul",
		Lastname:  "Minhaz",
		Address: &Address{
			City:  "Dhaka",
			State: "Dhaka",
		},
	})
	people = append(people, Person{
		ID:        "2",
		Firstname: "Ratul",
		Lastname:  "Minhaz",
	})
	people = append(people, Person{
		ID:        "3",
		Firstname: "Ratul",
		Lastname:  "Minhaz",
		Address: &Address{
			City:  "Dhaka",
			State: "Dhaka",
		},
	})

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":12345", router))
}
