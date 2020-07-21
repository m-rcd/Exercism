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

func Valid(number string) bool {
  formattedNumber := strings.ReplaceAll(number, " ", "")
	if len(formattedNumber) <= 1 { 
		return false 
	}
	
	if HasLetter(formattedNumber) {
		return false
	}

	var digits Ints
	var count = 0

	for i := len(formattedNumber) - 1;i >= 0; i-- {
		count += 1 
		integ,_ := strconv.Atoi(string(formattedNumber[i]))

		if count%2 == 0 {
		 sum := integ * 2 
		 if sum >= 9 {
			digits = append(digits, sum - 9)
     } else {
			digits = append(digits, sum)
		 }
		} else {
			digits = append(digits, integ)
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
