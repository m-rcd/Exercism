// Package etl transforms old score system to new system
package etl

import "strings"

func Transform(oldscore map[int][]string) map[string]int {

	NewScoreSystem := map[string]int{}

	for score, letters := range oldscore {
		for _, letter := range letters {
			NewScoreSystem[strings.ToLower(letter)] = score
		}
	}
	return NewScoreSystem
}
