package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
)

// JSONRequest ...
type JSONRequest struct {
	Action string `json:"action,omitempty"`
	Data   string `json:"data,omitempty"`
}

func debugRequest(r *http.Request) {
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(string(requestDump))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	base64decoder := base64.NewDecoder(base64.StdEncoding, r.Body)

	var rq JSONRequest
	_ = json.NewDecoder(base64decoder).Decode(&rq)

	fmt.Println("rq:", rq)
	fmt.Println("rq.Action:", rq.Action)
	fmt.Println("rq.Data:", rq.Data)
	fmt.Printf("Type: %T\n", rq)

	switch rq.Action {
	case "init-key-exchange":
		initKeyExchange()
	case "set-keys":
		setKeys()
	default:
		// fmt.Fprintln(w, "rq[action]:", rq.Action)
	}
}

func setKeys() {
	fmt.Println("set-keys")
}
func initKeyExchange() {
	fmt.Println("init-key-exchange")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handleRequest).Methods("POST")
	r.HandleFunc("/", hello)

	fmt.Println("HTTP Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
