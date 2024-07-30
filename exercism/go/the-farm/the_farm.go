package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	message string
}

// Error returns a string representation of the InvalidCowsError.
//
// No parameters.
// Returns a string.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// DivideFood calculates the amount of fodder to be divided among a given
// number of cows.
//
// Parameters: - calc: The FodderCalculator used to calculate the amount of
// fodder. - cows: The number of cows to divide the fodder among.
//
// Returns: - The amount of fodder to be divided among the cows. - An error if
// there was a problem calculating the fodder amount or fattening factor.
func DivideFood(calc FodderCalculator, cows int) (float64, error) {
	amount, err := calc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return amount * factor / float64(cows), nil
}

// ValidateInputAndDivideFood validates the input and divides the food among
// the cows.
//
// It takes a FodderCalculator object and the number of cows as parameters. It
// returns a float64 representing the divided food and an error if the number
// of cows is invalid.
func ValidateInputAndDivideFood(calc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(calc, cows)
}

// ValidateNumberOfCows validates the number of cows.
//
// It takes an integer parameter `cows` representing the number of cows.
// It returns an error type indicating any validation errors.
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{
			message: "there are no negative cows",
			cows:    cows,
		}
	case cows == 0:
		return &InvalidCowsError{
			message: "no cows don't need food",
			cows:    cows,
		}

	default:
		return nil
	}
}
