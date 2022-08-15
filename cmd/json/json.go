package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

const port = "8080"

func main() {

	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// encode a json response
// curl localhost:8080/encode
func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		First: "Jenny",
	}

	p2 := person{
		First: "James",
	}

	ps := []person{p1, p2}

	err := json.NewEncoder(w).Encode(ps)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

// decode a json request
// curl -H "Content-Type: application/json" -d'[{"First":"Jenny"},{"First":"James"}]' localhost:8080/decode
func decode(w http.ResponseWriter, r *http.Request) {
	ps := []person{}

	err := json.NewDecoder(r.Body).Decode(&ps)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Println(ps)
}
