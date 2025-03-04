# Simple go web server

- [Run and install the program](#How-to-run-and-install-the-program)
- [About this program](#About-this-program)
  - [Handler home](#Handler-home)
  - [handler about](#Handler-about)
  - [Handler contact](#Handler-contact)
  - [Handler json](#Handler-json)
  - [Static web page](#static-web-page)
  - [Go routine time](#Go-routine-time)
  - [Main function](#Main-function)


# How to run and install the program


## Install the Go:

You can install Go [here](https://go.dev/) or from your favorite package manager, depending on your operating system.


## First clone the repository:

```bash

git clone https://github.com/Helland369/GetPrepared.git

```


## Then cd in to the simple_web_server directory

```bash

cd ~/GetPrepared/golang/simple_web_server

```


## Then run the program:

```bash

go run main.go

```


# About this program

This is a simple web server written in Go. I made this as a school project to learn more about the Go programming language.


## Handler home

```go
func handler_home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the simple go web server!\n")
}
```

This function handles the home page, which is the first page you see when you access the server with curl or with a web browser. The text "welcome to the simple go web server!" is displayed. This function is connected to the HTTP server as shown in the [main function](#Main-function). You can see this page on: localhos:8080/


## Handler about

```go
func handler_about(w http.ResponseWriter, r *http.Request) {
	about := []string{"About this server:", "This is a very simple go web server", "This is very fun!\n"}
	for _, item := range about {
		fmt.Fprintln(w, item)
	}
}
```

This function handles the "about" page. This page displays a short text:
```
"About this server:
This is a very simple go web server
This is very fun!\n
```
You can see the HTTP server connection in the [main function](#Main-function) and access this page on: localhos:8080/about


## Handler contact

```go

func handler_contact(w http.ResponseWriter, r *http.Request) {
	contact := "This is the contact page. Currently there is no contact info, but you can try the yellow pages.\n"
	fmt.Fprintf(w, contact)
}

```

This function handles the contact page. The displayed text is:
```
This is the contact page. Currently there is no contact info, but you can try the yellow pages.
```
Since this is a school project I made the text humorous. If it hand been a real website there would have been proper information on this page. This page is found on: localhos:8080/contact , you can also see the HTTP server connection in the [main function](#Main-function).


## Handler json

This struct gets turned in to a json object:

```go

type person struct {
	Name string
	Age  int
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
		fmt.Fprintln(w, string(bs))
	}
}

```

This function returns a JSON object:

```
{"Name":"Ola","Age":20}
```

This page can be accessed on localhos:8080/json and you can see the connection in [main function](#Main-function).


## Static web page

```go

func static_web_page() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/static", http.StripPrefix("static", fileServer))

	fmt.Println("Static file server is running at http:://localhost:8080/static/page1.html")
}

```

This function handles a static HTML page. This page can be accessed on localhos:8080/static and you can see the connection in [main function](#Main-function).


## Go routine time

```go

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

```

This function runs in a separate goroutine and prints the current time to the terminal in 5 second intervals. It allows concurrency without interrupting the main server.


## Main function

The main function is where everything comes together and runs the web server. Here is a break down of each part.

First, the static web page handler and HTTP routes are set up:

```go
static_web_page()

mux := http.NewServeMux()
mux.HandleFunc("/", handler_home)
mux.HandleFunc("/about", handler_about)
mux.HandleFunc("/contact", handler_contact)
mux.HandleFunc("/json", handler_json)
```

Next the HTTP server configuration is initialized:

```go
server := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}
```

Channels are created to handle shutdown signals:

```go
stopChan := make(chan os.Signal, 1)
signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

goroutineStop := make(chan struct{})

go go_routine_time(goroutineStop)
```

The server starts running in a goroutine, allowing it to operate independently:

```go
go func() {
	fmt.Println("Starting server...")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("HTTP server listen and serve: %v\n", err)
	}
}()
```

When an input signal is received, the shutdown process begins:

```go
<-stopChan
fmt.Println("\nShutting down server...")

// stop the go routine
close(goroutineStop)
```

Finally, the server is gracefully shutdown, allowing any ongoing requests to complete:

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
    
if err := server.Shutdown(ctx); err != nil {
	fmt.Printf("Server is shuting down... %v\n", err)
}
fmt.Println("Server stoped!")
```

This structure ensures that the server runs smoothly and shuts down properly.

The whole main function:

```go
func main() {
	static_web_page()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler_home)
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
```
