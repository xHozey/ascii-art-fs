package web

import (
	"os"
	"strings"
)

// validates if the input contains only printable ASCII characters
func CheckValidInput(input string) string {
	filtredInput := ""
	for _, char := range input {
		if char == '\r' || char == '\n' {
			filtredInput += string(char)
		} else if int(char) < 32 || int(char) > 126 {
			continue
		} else {
			filtredInput += string(char)
		}
	}
	return filtredInput
}

func ReadBanner(banner string) map[rune][]string {
	data, _ := os.ReadFile("banners/" + banner + ".txt")
	stringData := string(data)
	if banner == "thinkertoy" {
		stringData = strings.ReplaceAll(stringData, "\r", "")
	}
	content := strings.Split(stringData[1:], "\n\n")
	characterMatrix := ConvertToCharacterMatrix(content)
	return characterMatrix
}
