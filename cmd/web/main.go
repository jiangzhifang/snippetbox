package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//log.Println("Starting server on :4000")
	log.Printf("Starting server on %s", *addr)
	//err := http.ListenAndServe(":4000", mux)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
