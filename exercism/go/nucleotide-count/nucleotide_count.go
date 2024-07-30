package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// newHistogram initializes and returns a new Histogram.
//
// No parameters.
// Returns a Histogram.
func newHistogram() Histogram {
	return Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts returns a histogram of the DNA sequence.
//
// It takes no parameters.
// It returns a Histogram, which is a map of runes to integers, and an error.
func (d DNA) Counts() (Histogram, error) {
	var histogram Histogram = newHistogram()
	for _, nucleotide := range d {
		switch nucleotide {
		case 'A', 'C', 'G', 'T':
			histogram[nucleotide]++
		default:
			return newHistogram(), errors.New("invalid nucleotide")
		}
	}
	return histogram, nil
}
