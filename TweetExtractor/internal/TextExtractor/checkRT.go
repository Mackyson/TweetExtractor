package TextExtractor

import (
	"strings"
)

func CheckBeRT(text string) bool {
	splitText := strings.Split(text, " ")
	return splitText[0] == "RT"
}
