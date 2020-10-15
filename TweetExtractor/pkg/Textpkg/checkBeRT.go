package Textpkg

import (
	"strings"
)

func CheckRT(text string) bool {
	splitText := strings.Split(text, " ")
	return splitText[0] == "RT"
}
