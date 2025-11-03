package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from glyst"))
}

func glystView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific Glyst with ID: %d", id)
	w.Write([]byte(msg))
}

func glystCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a specific Glyst"))
}

func glystCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /glyst/view/{id}", glystView)
	mux.HandleFunc("GET /glyst/create", glystCreate)
	mux.HandleFunc("POST /glyst/create", glystCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("starting server on: 4000")

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	fmt.Println(err)
	log.Fatal(err)
}
