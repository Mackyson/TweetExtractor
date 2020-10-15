package TextExtractor

import (
	// dbg "log"
	"regexp"
	"strings"
)

const (
	START_WORD = "at"
	END_WORD   = "in"
)

func ExtractSpotName(text string) string {
	var (
		startIdx = 0
		endIdx   = 0
	)
	splitText := strings.Split(text, " ")
	for i, s := range splitText {
		if s == START_WORD {
			startIdx = i
		}
		if s == END_WORD {
			endIdx = i
		}
		if endIdx == 0 && checkBeURL(s) {
			endIdx = i
		}
	}
	return strings.Join(splitText[startIdx+1:endIdx], " ")
}

func checkBeURL(text string) bool {
	return regexp.MustCompile(`http?`).Match([]byte(text))
}
