package main

import (
	"fmt"
	"net/http"
)

func main() {
	iconsDir := "./icons/"

	fs := http.FileServer(http.Dir(iconsDir))

	http.Handle("/icons/", http.StripPrefix("/icons/", fs))
	http.HandleFunc("/get/", serveIcons)
	http.HandleFunc("/post/centria", postIcons)

	fmt.Println("serving at port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
