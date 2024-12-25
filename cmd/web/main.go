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

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{

		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s \n", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

// Let's go file url: file:///Users/steeve/Documents/books/Alex%20Edwards%20-%20Let's%20Go%20(2022,%20Alex%20Edwards)%20-%20libgen.li.pdf
// Let's go file page number: 63
