package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Icon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	SIZE string `json:"size"`
}

func getIcons(folderPath string, searchTerm string, limit int) []Icon {
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
					newIcon.SIZE = getResolution(path)
					fmt.Println(path)
					fmt.Println(newIcon.SIZE)
					icons = append(icons, newIcon)
					if len(icons) > limit {
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

func getIconsFromArray(folderPath string, containsArr []string) []Icon {
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
				var newIcon Icon
				fileName := filepath.Base(path)
				fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
				if containsString(containsArr, fileNameWithoutExtension) {
					newIcon.Name = fileNameWithoutExtension
					newIcon.URL = "http://localhost:8080/" + path
					newIcon.SIZE = getResolution(path)
					fmt.Println(path)
					fmt.Println(newIcon.SIZE)
					icons = append(icons, newIcon)
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

func containsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}

func getResolution(imagePath string) string {

	if filepath.Ext(imagePath) == ".svg" {
		baseFolderName := getBaseFolderName(imagePath)
		if strings.Contains(baseFolderName, "@") {
			size := strings.Split(baseFolderName, "@")
			return size[0] + "x" + size[0]
		}
		return baseFolderName + "x" + baseFolderName
	}
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("Error opening image:", err)
		return "error"
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return "error"
	}
	if strings.Contains(imagePath, "@") {
		return strconv.Itoa(img.Width/2) + "x" + strconv.Itoa(img.Height/2)
	}
	return strconv.Itoa(img.Width) + "x" + strconv.Itoa(img.Height)
}
func getBaseFolderName(filePath string) string {
	dir := filepath.Dir(filePath)
	base := filepath.Base(dir)
	return base
}
