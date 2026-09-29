package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Asset struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go Asset API is running!")
}

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	assets := []Asset{
		{
			ID:        1,
			Name:      "Well A-01",
			Type:      "Oil Well",
			Latitude:  -2.1234,
			Longitude: 133.5678,
			Status:    "Active",
		},
		{
			ID:        2,
			Name:      "Well A-02",
			Type:      "Gas Well",
			Latitude:  -2.1567,
			Longitude: 133.5890,
			Status:    "Active",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(assets)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/assets", assetsHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
