package web

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	tpl, _ = template.ParseGlob("templates/*.html")
	banner := r.PostFormValue("banner")
	input := r.PostFormValue("input")
	if !CheckValidInput(input) {
		InvalidInput()
	}

	characterMatrix := ReadBanner(banner)
	s := DrawASCIIArt(characterMatrix, input)
	tpl.ExecuteTemplate(w, "index.html", s)
}
