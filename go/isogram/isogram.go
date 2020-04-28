package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	tracker := map[rune]int{}

	FormattedWord := strings.ReplaceAll(strings.ToLower(word), "-", "")
	FormattedWord = strings.ReplaceAll(FormattedWord, " ", "")

	for _, v := range FormattedWord {
		tracker[v] += 1
		if tracker[v] > 1 {
			return false
		}
	}
	return true
}
