package web

import (
	"fmt"
	"strings"
)

var Result string

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
func DrawASCIIArt(characterMatrix map[rune][]string, input string) string {
	Result := ""

	fmt.Println([]byte(input))

	splittedInput := strings.Split(input, "\r\n")

	for i, val := range splittedInput {
		if val == "" {
			if checkNewLine(splittedInput) {
				Result += "\n"
			} else if i != 0 && !checkNewLine(splittedInput) {
				Result += "\n"
			}
		} else if val != "" {
			for j := 0; j < 8; j++ {
				for _, k := range val {
					Result += characterMatrix[k][j]
				}
				Result += "\n"
			}
		}
	}
	fmt.Println(Result)
	return Result
}
