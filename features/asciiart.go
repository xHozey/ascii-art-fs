package web

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func AsciiArt(w http.ResponseWriter, r *http.Request) {
	tpl, _ = template.ParseGlob("templates/*.html")
	banner := r.PostFormValue("banner")
	fmt.Printf(banner)
	if banner == "" {
		banner = "standard"
		fmt.Println(banner)
	}
	input := r.PostFormValue("input")
	if !CheckValidInput(input) {
		InvalidInput()
	}

	// if r.Method == "GET" {
	// 	tpl.ExecuteTemplate(w, "index.html", nil)
	// } else if r.Method == "POST" {
	// 	http.Redirect(w, r, "/ascii-art", 200)
	// }

	characterMatrix := ReadBanner(banner)
	s := DrawASCIIArt(characterMatrix, input)
	tpl.ExecuteTemplate(w, "index.html", s)
}
