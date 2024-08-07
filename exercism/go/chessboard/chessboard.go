package chessboard

const FileCount = 8

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [FileCount]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if _, exists := cb[file]; !exists {
		return 0
	}

	count := 0
	for _, square := range cb[file] {
		if square {
			count++
		}
	}

	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	count := 0
	for _, file := range cb {
		if file[rank-1] {
			count++
		}
	}

	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * FileCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for name := range cb {
		count += CountInFile(cb, name)
	}

	return count
}
