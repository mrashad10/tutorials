package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps it takes to reach 1 using the Collatz Conjecture.
//
// It takes an integer parameter 'n' which represents the starting number.
// It returns an integer representing the number of steps taken, and an error if 'n' is less than 1.
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("only positive numbers are allowed")
	}

	steps := 0
	for n > 1 {
		switch n % 2 {
		case 0:
			n /= 2
		default:
			n = n*3 + 1
		}
		steps++
	}

	return steps, nil
}
