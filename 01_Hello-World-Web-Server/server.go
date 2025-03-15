package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful\n")
	name := r.FormValue("name")
	age := r.FormValue("age")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Age : %s", age)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server is Running at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
