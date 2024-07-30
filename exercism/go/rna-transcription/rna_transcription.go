package strand

import "strings"

var complements = map[string]string{
	"C": "G",
	"G": "C",
	"A": "U",
	"T": "A",
}

func ToRNA_old(dna string) string {
	var strand string
	for _, nucleotide := range dna {
		strand += complements[string(nucleotide)]
	}
	return strand
}

func ToRNA(dna string) string {
	var rnStrand strings.Builder
	for _, nucleotide := range dna {
		rnStrand.WriteString(complements[string(nucleotide)])
	}
	return rnStrand.String()
}

func ToRNA_replace(dna string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(dna)

}
