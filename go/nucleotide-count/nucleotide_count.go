package dna

import "errors"

type Histogram map[rune]int

type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	runes := []rune(d)
	for r := range runes {
		if _, ok := h[runes[r]]; !ok {
			return h, errors.New("invalid nucleotid")
		}
		h[runes[r]] += 1
	}

	return h, nil
}
