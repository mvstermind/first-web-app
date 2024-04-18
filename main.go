package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handler
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sup guys"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id")) // check if value of id is valid
	if err != nil {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("site id %d", id)
	w.Write([]byte(msg)) // display id number
}
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("site snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Println("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
