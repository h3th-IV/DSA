package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	items = append(items, Item{Id: 1, Name: "ProArt16", Quantity: 10})
	items = append(items, Item{Id: 1, Name: "ProArt17", Quantity: 15})

	fmt.Println("test server running at :7005")
	http.HandleFunc("/get-all", getAllItem)
	http.ListenAndServe(":7005", nil)
}

type Item struct {
	Id       int
	Name     string
	Quantity int
}

var items []Item

func getAllItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode("Method not allowed")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
