package grains

import (
	"errors"
	"math"
)

const q = 2

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("false")
	}
	return uint64(math.Pow(float64(q), float64(number-1))), nil
}

func Total() uint64 {
	var count uint64
	i := 1
	for 64 >= i {
		s, _ := Square(i)
		count += s
		i++
	}
	return count
}
