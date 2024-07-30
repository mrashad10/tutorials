package raindrops

import (
	"strconv"
	"strings"
)

// Factors returns the factors of a given number.
//
// The function takes an integer number as a parameter and returns a slice of
// integers.
func Factors(number int) []int {
	var factors []int

	for _, i := range []int{3, 5, 7} {
		if number%i == 0 {
			factors = append(factors, i)
		}
	}

	return factors
}

// Convert returns a string based on the factors of the given number.
//
// The function takes an integer parameter called 'number'.
// It returns a string.
func Convert(number int) string {
	var sounds strings.Builder
	factors := Factors(number)

	for _, factor := range factors {
		switch factor {
		case 3:
			sounds.WriteString("Pling")
		case 5:
			sounds.WriteString("Plang")
		case 7:
			sounds.WriteString("Plong")
		}
	}

	if sounds.Len() == 0 {
		return strconv.Itoa(number)
	}

	return sounds.String()
}
