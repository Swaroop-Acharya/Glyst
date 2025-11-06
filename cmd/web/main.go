package main

import (
	"flag"
	"log"
	"net/http"
)


type config struct{
	addr string
	staticDir string

}

func main() {

	var cfg config 

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	
	flag.StringVar(&cfg.addr, "addr",":4000","HTTP network address")
	flag.StringVar(&cfg.staticDir,"staticDir","./ui/static/","Path to static assests")
	flag.Parse()

	fileServer:=http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /glyst/view/{id}", glystView)
	mux.HandleFunc("GET /glyst/create", glystCreate)
	mux.HandleFunc("POST /glyst/create", glystCreatePost)

	// Print a log message to say that the server is starting.
	log.Print("starting server on:", cfg.addr )

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on (in this case ":4000")
	// and the servemux we just created. If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
