package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// SillyNephewError return error if cow are negative.
type SillyNephewError struct {
	count int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.count)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	w, err := weightFodder.FodderAmount()
	if cows == 0 {
		return 0, errors.New("division by zero")
	}
	if w < 0 {
		if err == nil {
			return 0, errors.New("negative fodder")
		}
		if err != ErrScaleMalfunction {
			return 0, err
		}
		return 0, errors.New("negative fodder")
	}
	if err == ErrScaleMalfunction {
		return 2 * w / float64(cows), nil
	}
	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}
	if err != nil {
		return 0, err
	}
	return w / float64(cows), err
}
