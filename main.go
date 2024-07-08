package main

import (
	"net/http"

	ft "web/features"
)

func main() {
	http.HandleFunc("/", ft.AsciiArt)
	http.ListenAndServe(":8081", nil)
}
