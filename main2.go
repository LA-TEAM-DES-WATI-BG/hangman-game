package main

import (
	"html/template"
	"net/http"
)

func main() {
	tmpl := template.Must(template.ParseFiles("hangman.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	http.ListenAndServe(":80", nil)
}
