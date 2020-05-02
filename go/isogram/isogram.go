package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	count := map[rune]int{}

	for _, letter := range word {
		if !unicode.IsLetter(letter) {
			continue
		}

		letter = unicode.ToLower(letter)
		count[letter] += 1
		if count[letter] > 1 {
			return false
		}

	}
	return true
}
