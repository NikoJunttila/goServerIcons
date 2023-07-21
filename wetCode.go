package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIconsCentria(w http.ResponseWriter, r *http.Request) {
	icons := getIcons("./icons/centria/", "")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewEncoder(w).Encode(icons)
	if err != nil {
		http.Error(w, "Couldn't encode JSON", http.StatusInternalServerError)
		return
	}
}
func serveIconsCentriaDark(w http.ResponseWriter, r *http.Request) {
	icons := getIcons("./icons/centria-dark/", "")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewEncoder(w).Encode(icons)
	if err != nil {
		http.Error(w, "Couldn't encode JSON", http.StatusInternalServerError)
		return
	}
}
func serveIconsBreeze(w http.ResponseWriter, r *http.Request) {
	icons := getIcons("./icons/breeze/", "")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	err := json.NewEncoder(w).Encode(icons)
	if err != nil {
		http.Error(w, "Couldn't encode JSON", http.StatusInternalServerError)
		return
	}
}
func serveIconsOxygen(w http.ResponseWriter, r *http.Request) {
	icons := getIcons("./icons/oxygen/", "")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
func postIconsCentria(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Parse the JSON payload from the request body
	var data IconData
	fmt.Printf(data.Theme)
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
	// You can access the received data now

	fmt.Printf("Icons: %v\n", data.Icons)
	// Sending the response back to the client as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArray)
}
func postIconsCentriaDark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	icons := getIcons("./icons/centria-dark/", "")
	neededIcons := parseItems(string(body))
	var newArray []Icon
	for _, icon := range icons {
		if containsString(neededIcons, icon.Name) {
			newArray = append(newArray, icon)
		}
	}

	// Send the response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArray)
}
func postIconsBreeze(w http.ResponseWriter, r *http.Request) {
	fmt.Println("connection ...")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	icons := getIcons("./icons/breeze/", "")
	neededIcons := parseItems(string(body))

	var newArray []Icon
	for _, icon := range icons {
		if containsString(neededIcons, icon.Name) {
			newArray = append(newArray, icon)
		}
	}
	fmt.Println("sending....")
	// Send the response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArray)
}
func postIconsOxygen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	icons := getIcons("./icons/oxygen/", "")
	neededIcons := parseItems(string(body))
	var newArray []Icon
	for _, icon := range icons {
		if containsString(neededIcons, icon.Name) {
			newArray = append(newArray, icon)
		}
	}

	// Send the response back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newArray)
}
