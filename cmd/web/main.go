package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// create a file server which serves files out of the './ui/static'
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// use the mux.Handle() function to register the file server as the handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
