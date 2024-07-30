package resistorcolor

// colorsValues is a map that associates each color band on a resistor with
// its corresponding numerical value. The keys are the names of the colors in
// lowercase, and the values are the numerical values of the color bands.
var colorsValues = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors returns a slice of strings containing all the colors available.
//
// No parameters.
// Returns a slice of strings.
func Colors() []string {
	result := make([]string, 0, len(colorsValues))
	for key := range colorsValues {
		result = append(result, key)
	}
	return result
}

// ColorCode returns the corresponding integer code for a given color.
//
// Parameters:
// - color: a string representing a color.
//
// Return type:
// - int: the integer code corresponding to the color.
func ColorCode(color string) int {
	return colorsValues[color]
}
