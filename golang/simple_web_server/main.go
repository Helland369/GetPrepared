package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the simple go web server!")
}

func handler_about(w http.ResponseWriter, r *http.Request) {
	about := []string{"About this server:", "This is a very simple go web server", "This is very fun!"}

	for _, item := range about {
		fmt.Fprintln(w, item)
	}
}

func handler_contact(w http.ResponseWriter, r *http.Request) {
contact := "Prøv oplysningen!"
fmt.Fprintf(w, contact)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", handler_about)
	http.HandleFunc("/contact", handler_contact)

	http.ListenAndServe(":8080", nil)
}
