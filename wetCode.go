package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

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

	icons := getIcons(fmt.Sprintf("./icons/%s/", selectedValue), "")
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
func postIcons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse the JSON payload from the request body
	var data IconData
	err := json.NewDecoder(r.Body).Decode(&data)
	icons := getIcons(fmt.Sprintf("./icons/%s/", data.Theme), "")
	var newArray []Icon
	if err != nil {
		fmt.Println("error oopsie")
	}
	for _, icon := range icons {
		if containsString(data.Icons, icon.Name) {
			newArray = append(newArray, icon)
		}
	}
	// Sending the response back to the client as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArray)
}
