package TextExtractor

import (
	"strings"
)

func CheckWhetherRT(text string) bool {
	splitText := strings.Split(text, " ")
	return splitText[0] == "RT"
}
