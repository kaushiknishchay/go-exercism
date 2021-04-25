package grains

import (
	"errors"
)

func Square(squareNum int) (uint64, error) {
	if squareNum <= 0 || squareNum > 64 {
		return 0, errors.New("square number has to be between 1 & 64")
	}

	return 1 << (squareNum - 1), nil
}

func Total() uint64 {
	total, err := Square(64)

	if err != nil {
		panic("failed to calculate total")
	}

	return total*2 - 1
}
