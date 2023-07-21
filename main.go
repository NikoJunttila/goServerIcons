package main

import (
	"fmt"
	"net/http"
)

func main() {
	iconsDir := "./icons/"

	fs := http.FileServer(http.Dir(iconsDir))

	http.Handle("/icons/", http.StripPrefix("/icons/", fs))
	http.HandleFunc("/icons/centria", serveIconsBreeze)
	http.HandleFunc("/icons/centriaDark", serveIconsCentriaDark)
	http.HandleFunc("/icons/breeze", serveIconsBreeze)
	http.HandleFunc("/icons/oxygen", serveIconsOxygen)
	http.HandleFunc("/post/centria", postIconsCentria)
	http.HandleFunc("/post/centriaDark", postIconsCentriaDark)
	http.HandleFunc("/post/breeze", postIconsBreeze)
	http.HandleFunc("/post/oxygen", postIconsOxygen)

	fmt.Println("serving at port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
