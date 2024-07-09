package web

import (
	"html/template"
	"net/http"
)

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index1.html")
	if err != nil {
		http.Error(w, "hello", 500)
	}
	r.ParseForm()
	banner := r.Form.Get("banner")
	input := r.Form.Get("input")

	filtredInput := CheckValidInput(input)

	characterMatrix := ReadBanner(banner)
	s := DrawASCIIArt(characterMatrix, filtredInput)

	tpl.Execute(w, s)
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "hello", 500)
	}

	tpl.Execute(w, nil)
}
