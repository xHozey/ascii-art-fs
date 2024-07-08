package web

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	tpl, _ = template.ParseGlob("templates/*.html")
	banner := r.FormValue("banner")
	fmt.Print(banner)
	input := r.FormValue("input")
	if !CheckValidInput(input) {
		InvalidInput()
	}

	characterMatrix := ReadBanner(banner)
	s := DrawASCIIArt(characterMatrix, input)
	tpl.ExecuteTemplate(w, "index.html", s)
}
