package reverse

func Reverse(s string) string { 
	reversedString := ""
	runes := []rune(s)
  
	for i := len(runes) - 1; i >= 0; i-- {
		reversedString += string(runes[i])
	}

	return reversedString
}