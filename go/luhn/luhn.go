package luhn

import (
    "strconv"
		"strings"
		"unicode"
)

// Checks whether a number is valid or not based on luhn algorithm
func Valid(number string) bool {
  formattedNumber := strings.ReplaceAll(number, " ", "")
	if len(formattedNumber) <= 1  { 
		return false 
	}

	for _, r := range formattedNumber {
		if unicode.IsLetter(r) {
			return false
	  }	
	}
	var count int
	double := len(formattedNumber)%2 == 0 
	
	for _, r := range formattedNumber {
		digit,_ := strconv.Atoi(string(r))

		if double && digit * 2  >= 9 {
			count +=  digit * 2 - 9
		}
		if double && digit * 2 < 9 {
				count += digit * 2
		}
		if !double {
			count += digit
		}
		double = !double
	}
	return count%10==0
}
