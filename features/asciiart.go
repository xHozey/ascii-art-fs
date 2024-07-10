package web

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/index.html")

	tpl.Execute(w, nil)
}

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/index1.html")

	r.ParseForm()
	banner := r.Form.Get("banner")
	input := r.Form.Get("input")

	filtredInput := CheckValidInput(input)

	characterMatrix := ReadBanner(banner)
	s := DrawASCIIArt(characterMatrix, filtredInput)

	tpl.Execute(w, s)
}
