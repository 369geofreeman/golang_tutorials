package grains

import (
	"errors"
)

// Returns number, err; how many grains were on a given square
func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("error")
	}

	total := 1
	for i := 1; i < number; i++ {
		total += total
	}

	return uint64(total), nil
}

// Returns the total number of grains on the chessboard
func Total() uint64 {
	var ret uint64
	var err error

	for i := 1; i < 65; i++ {
		num, e := Square(i)
		ret += num
		err = e
	}

	if err == nil {
		return ret
	}
	return 0
}
