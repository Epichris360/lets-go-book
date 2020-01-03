package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// page 115 - database time
func main() {
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	// creates a new logger. standard out, extra text, additional info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// initialize new http.Server struct to carry the error log func, port and mux
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on port%s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
