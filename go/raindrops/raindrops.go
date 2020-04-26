package raindrops

import "strconv"

const (
	THREE = 3
	FIVE  = 5
	SEVEN = 7
)

type rule struct {
	divisor int
	sound   string
}

func Convert(number int) string {
	rules := []rule{
		{THREE, "Pling"},
		{FIVE, "Plang"},
		{SEVEN, "Plong"},
	}

	word := ""
	for _, r := range rules {
		if number%r.divisor == 0 {
			word += r.sound
		}
	}

	if word == "" {
		word += strconv.Itoa(number)
	}
	return word
}
