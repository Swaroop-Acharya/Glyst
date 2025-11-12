package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {

	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /glyst/view/{id}", app.glystView)
	mux.HandleFunc("GET /glyst/create", app.glystCreate)
	mux.HandleFunc("POST /glyst/create", app.glystCreatePost)
	return app.recoverPanic(app.logRequest(commonHeader(mux)))
}
