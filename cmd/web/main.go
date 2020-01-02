package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// page 82
func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// creates a new logger. standard out, extra text, additional info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// create a file server which serves files out of the './ui/static'
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// use the mux.Handle() function to register the file server as the handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// initialize new http.Server struct to carry the error log func, port and mux
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on port%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
