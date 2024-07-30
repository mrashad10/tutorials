package hamming

import "errors"

// Distance calculates the difference between two given strings.
//
// It takes two strings, `a` and `b`, as input parameters. Both strings must
// be of equal length.
//
// It returns an integer that represents the number of positions where the
// characters in the two strings differ, and an error if the lengths of the
// strings are not equal.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("strands must be of equal length")
	}

	diff := 0
	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
