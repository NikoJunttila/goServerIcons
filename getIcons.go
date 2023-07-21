package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Icon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getIcons(folderPath string, searchTerm string) []Icon {
	var icons []Icon

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if strings.HasSuffix(strings.ToLower(info.Name()), ".svg") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".png") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".jpg") ||
				strings.HasSuffix(strings.ToLower(info.Name()), ".jpeg") {
				if searchTerm == "" || strings.Contains(strings.ToLower(info.Name()), strings.ToLower(searchTerm)) {
					var newIcon Icon
					fileName := filepath.Base(path)
					fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
					newIcon.Name = fileNameWithoutExtension
					newIcon.URL = "http://localhost:8080/" + path
					icons = append(icons, newIcon)
					if len(icons) >= 200 {
						return filepath.SkipDir
					}
				}
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	return icons
}
func parseItems(inputString string) []string {
	// Step 1: Remove the brackets and quote marks from the input string
	inputString = strings.ReplaceAll(inputString, "[", "")
	inputString = strings.ReplaceAll(inputString, "]", "")
	inputString = strings.ReplaceAll(inputString, `"`, "")

	// Step 2: Split the input string based on comma separator
	items := strings.Split(inputString, ",")

	// Step 3: Create a new array to hold the cleaned-up items
	var cleanedItems []string

	// Step 4: Loop through each item to clean it up and add it to the cleanedItems array
	for _, item := range items {
		cleanedItem := strings.TrimSpace(item) // Remove leading and trailing whitespaces
		cleanedItems = append(cleanedItems, cleanedItem)
	}

	return cleanedItems
}
func containsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
