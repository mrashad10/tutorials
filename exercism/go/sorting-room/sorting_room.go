package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox returns a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

// Value returns the value of the FancyNumber object as a string.
//
// No parameters.
// Returns a string.
func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber returns the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	// Check if fnb is a FancyNumber
	fn, ok := fnb.(FancyNumber)
	if !ok {
		return 0
	}

	// Get the value from FancyNumber
	value := fn.Value()

	// Convert the value to an integer
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return intValue
}

// DescribeFancyNumberBox efficiently describes the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	// Check if fnb is of type FancyNumber
	if _, ok := fnb.(FancyNumber); ok {
		// If fnb is a FancyNumber, extract the number and format the string
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
	}

	// If fnb is not a FancyNumber, return default string
	return "This is a fancy box containing the number 0.0"
}

// DescribeAnything returns a string describing the given input.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
