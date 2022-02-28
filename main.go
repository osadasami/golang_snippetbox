package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
