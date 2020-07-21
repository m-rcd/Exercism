package luhn

import (
    "strconv"
		"strings"
		"unicode"
)

type Ints []int

func HasLetter(word string) bool {
	for _, r := range word {
		if unicode.IsLetter(r) {
			return true
	  }	
	}
	return false
}

// Checks whether a number is valid or not based on luhn algorithm
func Valid(number string) bool {
  formattedNumber := strings.ReplaceAll(number, " ", "")
	if len(formattedNumber) <= 1  || HasLetter(formattedNumber) { 
		return false 
	}

	var digits Ints
	var count = 0

	for i := len(formattedNumber) - 1;i >= 0; i-- {
		count += 1 
		digit,_ := strconv.Atoi(string(formattedNumber[i]))

		if count%2 == 0 {
		 double := digit * 2 
		 if double >= 9 {
			digits = append(digits, double - 9)
     } else {
			digits = append(digits, double)
		 }
		} else {
			digits = append(digits, digit)
	  }
	}

	var digitSum int 
	for _, r := range digits {
		digitSum += r 
	}
	if digitSum%10 ==0 {
		return true
	} else {
		return false
	}
}
