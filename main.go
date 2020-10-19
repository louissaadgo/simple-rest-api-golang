package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type userModel struct {
	Name       string `json:"name"`
	FamilyName string `json:"familyName"`
	Age        int    `json:"age"`
}

var users = []userModel{}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func addUsers(w http.ResponseWriter, r *http.Request) {
	var user userModel
	json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	fmt.Fprint(w, "Received User :)")
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", addUsers).Methods("POST")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func main() {
	handleRequests()
}
