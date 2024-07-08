package main

import (
	"fmt"
	"html/template"
	"net/http"

	ft "web/features"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, r)
}

func main() {
	ft.AsciiArt()
	http.HandleFunc("/", serveStatic)
	http.ListenAndServe(":8080", nil)
}
