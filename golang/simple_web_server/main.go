package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type person struct {
	Name string
	Age  int
}

// localhos:8080/
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the simple go web server!\n")
}

// localhos:8080/about
func handler_about(w http.ResponseWriter, r *http.Request) {
	about := []string{"About this server:", "This is a very simple go web server", "This is very fun!\n"}
	for _, item := range about {
		fmt.Fprintln(w, item)
	}
}

// localhos:8080/contact
func handler_contact(w http.ResponseWriter, r *http.Request) {
	contact := "This is the contact page. Currently there is no contact info, but you can try the yellow pages.\n"
	fmt.Fprintf(w, contact)
}

// localhos:8080/json
func handler_json(w http.ResponseWriter, r *http.Request) {
	newPerson := person{
		Name: "Ola",
		Age:  20,
	}
	bs, err := json.Marshal(newPerson)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintln(w, string(bs))
	}
}

// localhos:8080/static
func static_web_page() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static", http.StripPrefix("static", fileServer))

	fmt.Println("Static file server is running at http:://localhost:8080/static/page1.html")
}

// prints the current time every 5 sec in the terminal
func go_routine_time(stopChan chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Time: ", time.Now().Format("15:04:05"))
		case <-stopChan:
			fmt.Println("Stopping time go routine...")
			return
		}
	}
}

func main() {
	static_web_page()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/about", handler_about)
	mux.HandleFunc("/contact", handler_contact)
	mux.HandleFunc("/json", handler_json)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// chanel for handeling shutdown signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// chanel for stoping go_routine_time
	goroutineStop := make(chan struct{})

	go go_routine_time(goroutineStop)

	// start server in a go routine
	go func() {
		fmt.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("HTTP server listen and serve: %v\n", err)
		}
	}()

	// wait for interupt signal
	<-stopChan
	fmt.Println("\nShutting down server...")

	// stop the go routine
	close(goroutineStop)

	// Create context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful server shutdown
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Server is shuting down... %v\n", err)
	}
	fmt.Println("Server stoped!")
}
