package asciiart

import (
	"fmt"
	"strings"
)

// Convert content array to a character matrix mapping ASCII characters to their line representations
func ConvertToCharacterMatrix(content []string) map[rune][]string {
	characterMatrix := map[rune][]string{}
	for i, val := range content {
		characterMatrix[rune(32+i)] = strings.Split(val, "\n")
	}
	return characterMatrix
}

// Check if there are any non-empty lines in the input lines array
func checkNewLine(splittedInput []string) bool {
	var hasNonEmptyLines bool
	for _, line := range splittedInput {
		if line != "" {
			hasNonEmptyLines = true
			break
		}
	}
	return hasNonEmptyLines
}

// Render the ASCII art based on the character matrix and the input lines
func DrawASCIIArt(characterMatrix map[rune][]string, input string) {

	splittedInput := strings.Split(input, "\\n")

	for i, val := range splittedInput {
		if val == "" {
			if checkNewLine(splittedInput) {
				fmt.Println()
			} else if i != 0 && !checkNewLine(splittedInput) {
				fmt.Println()
			}
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {

					fmt.Print(characterMatrix[k][j])
				}
				fmt.Println()
			}
		}
	}
}
