package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	message string
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("%s", e.message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	if cows < 0 {
		return 0.0, &SillyNephewError{message: fmt.Sprintf("silly nephew, there cannot be %d cows", cows)}
	}

	fodder, err := WeightFodder.FodderAmount(weightFodder)

	if err == nil && fodder >= 0 && cows > 0 {
		c := float64(cows)
		return fodder / c, err
	}

	if err == ErrScaleMalfunction && fodder > 0 && cows > 0 {
		return fodder / float64(cows) * 2, nil
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if fodder < 0 && err == ErrScaleMalfunction || err == nil {
		return 0.0, errors.New("negative fodder")
	}
	return 0, err
}
