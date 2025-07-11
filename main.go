package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mult := mux.NewRouter()
	mult.HandleFunc("/get-item", GetAnmItem).Methods(http.MethodGet)
	mult.HandleFunc("/new-item", addNewItem).Methods(http.MethodPost)
	mult.HandleFunc("/health", CheckServerHealth).Methods((http.MethodGet))

	fmt.Println("server running at port 9000")
	http.ListenAndServe(":9000", mult)

}

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Size string `json:"size"`
}

func addNewItem(w http.ResponseWriter, r *http.Request) {
	var item Item

	w.Header().Add("Content-type", "application/json")
	err := json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("unable to process request body")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func CheckServerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Server Connected Successfully")
}

func GetAnmItem(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode("Item fetched successfully")
}
