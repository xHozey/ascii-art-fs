package main

import (
	"os"

	ft "asciiart/features"
)

func main() {
	args := os.Args[1:]

	// Specify the ASCII art banner file to use
	input, banner := ft.BannerInput(args)

	if !ft.CheckValidInput(input) {
		ft.InvalidInput()
	}
	characterMatrix := ft.ReadBanner(banner)
	ft.DrawASCIIArt(characterMatrix, input)
}
