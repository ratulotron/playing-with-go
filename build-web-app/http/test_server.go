package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", v)
	}

	fmt.Fprintf(w, "Hello world!")
}

//ShowAllTasksFunc is used to handle the "/" URL which is the default ons
//TODO add http404 error
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// context := db.GetTasks("pending") //true when you want non deleted tasks
		// //db is a package which interacts with the database
		// if message != "" {
		// 	context.Message = message
		// }
		// homeTemplate.Execute(w, context)
		fmt.Fprintf(w, "You GOT me!")
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

//DeleteTaskFunc is used to delete a task, trash = move to recycle bin, delete = permanent delete
func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := r.URL.Path[len("/delete/"):]
		if id == "all" {
			fmt.Fprintf(w, "Are you seriously gonna delete all?!")
		} else {
			fmt.Fprintf(w, "Are you seriously gonna delete %s?!", id)
		}
	} else {
		fmt.Fprintf(w, "Why are you here?!")
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/taskl/", ShowAllTasksFunc)
	// http.HandleFunc("/complete/", CompleteTaskFunc)
	http.HandleFunc("/delete/", DeleteTaskFunc)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
