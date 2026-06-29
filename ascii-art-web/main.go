package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleAscii(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Result string
	}

	if r.Method != "POST" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if banner != "standard.txt" && banner != "shadow.txt" && banner != "thinkertoy.txt" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	charMap := loadBanner(banner)
	ascii := generateAscii(text, charMap)
	data := PageData{Result: ascii}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

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

func loadBanner(fileName string) map[rune][]string {

	var readByte []byte
	var err error

	switch fileName {
	case "shadow.txt":
		readByte, err = os.ReadFile("shadow.txt")
	case "thinkertoy.txt":
		readByte, err = os.ReadFile("thinkertoy.txt")
	default:
		readByte, err = (os.ReadFile("standard.txt"))
	}
	errorHandling(err)

	toString := string(readByte)
	toString = strings.ReplaceAll(toString, "\r\n", "\n")
	splitStr := strings.Split(toString, "\n")

	charMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		startLine := (i - 32) * 9
		charLines := splitStr[startLine+1 : startLine+9]
		charMap[rune(i)] = charLines
	}

	return charMap
}

func generateAscii(args string, charMap map[rune][]string) string {
	var result string
	lines := strings.Split(args, "\\n")
	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				result += "\n"
			}
			continue
		}

		row := make([]string, 8)

		for _, char := range line {
			for i := 0; i < 8; i++ {
				row[i] += charMap[char][i]
			}
		}
		for _, r := range row {
			result += r + "\n"
		}
	}
	return result
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ascii-art", handleAscii)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
