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
		//Handler:  mux,
		Handler: app.routes(),
	}

	//log.Println("Starting server on :4000")
	//log.Printf("Starting server on %s", *addr)
	infoLog.Printf("Starting server on %s", *addr)
	//err := http.ListenAndServe(":4000", mux)
	err := srv.ListenAndServe()
	//log.Fatal(err)
	errorLog.Fatal(err)
}
