package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	iconsDir := "./icons/"

	fs := http.FileServer(http.Dir(iconsDir))

	http.Handle("/icons/", http.StripPrefix("/icons/", fs))
	http.HandleFunc("/get/", serveIcons)
	http.HandleFunc("/post/centria", postIcons)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("serving at port: 8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
