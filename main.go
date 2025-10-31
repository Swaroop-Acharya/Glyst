package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from glyst"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("starting server on: 4000")

	err := http.ListenAndServe(":4000", mux)
	fmt.Println(err)
	log.Fatal(err)
}
