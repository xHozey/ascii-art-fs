package asciiart

import (
	"fmt"
	"os"
)

func Usage() {
	fmt.Printf("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
	os.Exit(0)
}

func InvalidInput() {
	fmt.Println("Error: The input contains characters without corresponding ASCII art representation!")
	os.Exit(0)
}

func InvalidBanner() {
	fmt.Printf("invalid bannerfile name")
	os.Exit(0)
}
