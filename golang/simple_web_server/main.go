package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type person struct {
	Name string
	Age  int
}

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

func handler_json(w http.ResponseWriter, r *http.Request) {
	newPerson := person{
		Name: "Ola",
		Age:  20,
	}
	bs, err := json.Marshal(newPerson)
	if err != nil {
		panic(err)
	} else {
		//fmt.Fprintln(w, bs)
		fmt.Fprintln(w, string(bs))
	}
}

func static_web_page() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static", http.StripPrefix("static", fileServer))

	fmt.Println("Static file server is running at http:://localhost:8080/static/page1.html")
}

func go_routine_time() {
	for {
		fmt.Println("Time: ", time.Now().Format("15:04:05"))
		time.Sleep(5 * time.Second)
	}
}

func main() {
	static_web_page()

	http.HandleFunc("/", handler)
	http.HandleFunc("/about", handler_about)
	http.HandleFunc("/contact", handler_contact)
	http.HandleFunc("/json", handler_json)

	go go_routine_time()

	http.ListenAndServe(":8080", nil)
}
