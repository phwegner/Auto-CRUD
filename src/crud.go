package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Id          string `json:"id"`
	Data_string string `json:"data"`
}

var dat []Data

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	dat = append(dat, data)
	json.NewEncoder(w).Encode(dat)
}

func get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range dat {
		if item.Id == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - No such id"))
}

func put(w http.ResponseWriter, r *http.Request) {
	var data Data
	_ = json.NewDecoder(r.Body).Decode(&data)
	fmt.Println(data)
	for _, item := range dat {
		if item.Id == data.Id {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 - No such id"))

}

// Existing code from above
func handleRequests() {
	// creates a new instance of a mux router
	router := mux.NewRouter()
	// replace http.HandleFunc with myRouter.HandleFunc
	router.HandleFunc("/create", create).Methods("POST")
	router.HandleFunc("/get/{id}", get).Methods("GET")
	router.HandleFunc("/put", put).Methods("PUT", "POST")

	log.Fatal(http.ListenAndServe(":9009", router))
}

func main() {
	fmt.Println("Rest API v1.0 - ##--Auto-CRUD--##")
	//mock data
	dat = append(dat, Data{Id: "2", Data_string: "test"})

	// Start server
	handleRequests()
}
