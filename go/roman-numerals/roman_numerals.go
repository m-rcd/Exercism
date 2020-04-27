package romannumerals

import (
	"errors"
)

func ToRomanNumeral(number int) (string, error) {
	arabic := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	if number <= 0 || number > 3000 {
		return result, errors.New("invalid number")
	}

	for i, v := range arabic {
		for number >= v {
			result += numerals[i]
			number -= v
		}
	}

	return result, nil
}
