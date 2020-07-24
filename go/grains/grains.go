package grains

import "errors"

// Square returns how many grains there are in the square provided
func Square(input int) (uint64, error) {
	if input <  1 || input > 64 {
		return uint64(input), errors.New("Invalid input")
	}
	var count []uint64 
	count = append(count, 1) 

	for i := 1; i <= input - 1; i++ {
		count = append(count, count[len(count)-1] * 2)
	}
	return count[len(count)-1], nil
}

// total returns the number of grains on the
func Total() uint64 {
	var count []uint64
	count = append(count, 1) 

	for i := 1; i <= 64; i++ {
		count = append(count, count[len(count)-1] * 2)
	}
	return count[len(count)-1]
}