package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Strands have different sizes!")
	}

	count := 0
	runesA := []rune(a)
	runesB := []rune(b)
	for i := range runesA {
		if runesA[i] != runesB[i] {
			count += 1
		}
	}
	return count, nil
}
