package main

import (
	"fmt"
	"net/http"

	ft "web/features"
)

func main() {
	http.HandleFunc("/", ft.AsciiArt)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print(err)
	}
}
