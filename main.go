package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my world!"))

}

func main() {
	mux := http.NewServeMux() //need to setup router to register handler
	mux.HandleFunc("/", home)
	log.Println("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
