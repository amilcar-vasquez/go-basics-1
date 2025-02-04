package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my world!"))

}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

func main() {
	mux := http.NewServeMux() //need to setup router to register handler
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
