package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ASCII Art output here")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art", handleHome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
