package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type JsonRequest struct {
	action string
	data   string
}

func doStuff(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")

	fmt.Printf("%v", r.Body)

	var rq JsonRequest
	_ = json.NewDecoder(r.Body).Decode(&rq)

	fmt.Println("rq:", rq)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", doStuff).Methods("POST")
	r.HandleFunc("/", doStuff)

	fmt.Println("HTTP Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
