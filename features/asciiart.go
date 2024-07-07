package web

import (
	"os"
)

func AsciiArt() {
	args := os.Args[1:]

	// Specify the ASCII art banner file to use
	input, banner := BannerInput(args)

	if !CheckValidInput(input) {
		InvalidInput()
	}
	characterMatrix := ReadBanner(banner)
	DrawASCIIArt(characterMatrix, input)
}
