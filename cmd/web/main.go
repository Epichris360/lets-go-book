package main

import (
	"flag"
	"log"
	"net/http"
)

// page 82
func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// create a file server which serves files out of the './ui/static'
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// use the mux.Handle() function to register the file server as the handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on port%s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
