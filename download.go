package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func handleDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var data IconData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	// Provide the folder path containing the icons (adjust it according to your setup)
	iconFolder := fmt.Sprintf("./icons/%s/", data.Theme)
	fmt.Println(data.Icons)
	// Create a buffer to store the zip file content
	zipBuffer := new(bytes.Buffer)

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipBuffer)

	// Walk through the folder to filter and add selected icons to the zip file
	filepath.Walk(iconFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileName := filepath.Base(path)
		fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		if !info.IsDir() && containsString(data.Icons, fileNameWithoutExtension) {
			fmt.Println(fileNameWithoutExtension)
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			relativePath, err := filepath.Rel(iconFolder, path)
			if err != nil {
				return err
			}

			zipFile, err := zipWriter.Create(relativePath)
			if err != nil {
				return err
			}
			_, err = io.Copy(zipFile, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	err = zipWriter.Close()
	if err != nil {
		http.Error(w, "Error creating the zip file", http.StatusInternalServerError)
		return
	}
	fmt.Println("sent response")
	// Set the appropriate headers for the download
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=icons.zip")

	// Write the zip file content to the response
	_, err = zipBuffer.WriteTo(w)
	if err != nil {
		http.Error(w, "Error sending the zip file", http.StatusInternalServerError)
		return
	}
}
