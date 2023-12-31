package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// get request
func serveIcons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	path := r.URL.Path
	segments := strings.Split(path, "/")

	var selectedValue string
	for i, seg := range segments {
		if seg == "get" && i+1 < len(segments) {
			selectedValue = segments[i+1]
			break
		}
	}

	domain := r.Host
	icons := getIcons(fmt.Sprintf("./icons/%s/", selectedValue), "", 40, domain)
	err := json.NewEncoder(w).Encode(icons)
	if err != nil {
		http.Error(w, "Couldn't encode JSON", http.StatusInternalServerError)
		return
	}
}

type IconData struct {
	Theme string   `json:"theme"`
	Icons []string `json:"icons"`
}

// POSTS
func iconsFromListOfNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse the JSON payload from the request body
	var data IconData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
  domain := r.Host
	icons := getIconsFromArray(fmt.Sprintf("./icons/%s/", data.Theme), data.Icons,domain)
	// Sending the response back to the client as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(icons)
}

type searchReq struct {
	Theme      string `json:"theme"`
	SearchTerm string `json:"searchTerm"`
}

func searchIcons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse the JSON payload from the request body
	var data searchReq
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println("error oopsie")
	}
	domain := r.Host
	icons := getIcons(fmt.Sprintf("./icons/%s/", data.Theme), data.SearchTerm, 1000, domain)
	// Sending the response back to the client as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(icons)
}
