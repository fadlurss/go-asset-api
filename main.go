package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Asset struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
}

var assets = []Asset{
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

func assetsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path == "/assets" {
		json.NewEncoder(w).Encode(assets)
		return
	}

	// /assets/1
	idText := strings.TrimPrefix(r.URL.Path, "/assets/")
	id, err := strconv.Atoi(idText)

	if err != nil {
		http.Error(w, `{"error":"Invalid asset ID"}`, http.StatusBadRequest)
		return
	}

	for _, asset := range assets {
		if asset.ID == id {
			json.NewEncoder(w).Encode(asset)
			return
		}
	}

	http.Error(w, `{"error":"Asset not found"}`, http.StatusNotFound)
}

func main() {
	http.HandleFunc("/assets/", assetsHandler)
	http.HandleFunc("/assets", assetsHandler)

	println("Go Asset API is running!")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		println("Server error:", err.Error())
	}
}
