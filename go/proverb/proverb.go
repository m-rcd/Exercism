// package proverb will return a proverb using inputs words
package proverb

import "fmt"

// Proverb returns a proverb
func Proverb(rhyme []string) []string {
	lastLine := "And all for the want of a %s."
	if len(rhyme) == 1 {
		return []string{fmt.Sprintf(lastLine, rhyme[0])}
	}
	proverb := []string{}
	for i := 0; i < len(rhyme); i++ {
		if i < len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
		if i == len(rhyme)-1 {
			proverb = append(proverb, fmt.Sprintf(lastLine, rhyme[0]))
		}
	}

	return proverb
}
