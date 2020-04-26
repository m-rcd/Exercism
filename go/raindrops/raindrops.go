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
	rule3 := rule{divisor: THREE, sound: "Pling"}
	rule5 := rule{divisor: FIVE, sound: "Plang"}
	rule7 := rule{divisor: SEVEN, sound: "Plong"}

	rules := []rule{rule3, rule5, rule7}

	word := ""
	for r := range rules {
		if number%rules[r].divisor == 0 {
			word += rules[r].sound
		}
	}

	if word == "" {
		word += strconv.Itoa(number)
	}
	return word
}
