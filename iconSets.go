package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func getIconSetsNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	iconsDir := "./icons/"
	subfolders, err := listSubfolders(iconsDir)
	if err != nil {
		fmt.Printf("Error listing subfolders: %s\n", err)
	}

	err = json.NewEncoder(w).Encode(subfolders)
	if err != nil {
		http.Error(w, "Couldn't encode JSON", http.StatusInternalServerError)
		return
	}
}

func listSubfolders(folderPath string) ([]string, error) {
	var subfolders []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != folderPath {
			subfolder := filepath.Base(path)
			subfolders = append(subfolders, subfolder)
			return filepath.SkipDir // To prevent walking inside subfolders
		}

		return nil
	})

	return subfolders, err
}
