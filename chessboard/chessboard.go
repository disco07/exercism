package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var sum int
	for _, v := range cb[rank] {
		if v {
			sum++
		}
	}
	return sum
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var sum int
	if file < 1 || file > 8 {
		return sum
	}
	for i, _ := range cb {
		if cb[i][file-1] {
			sum++
		}
	}
	return sum
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var sum int
	for _, v := range cb {
		for _, _ = range v {
			sum++
		}
	}
	return sum
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var sum int
	for _, rank := range cb {
		for _, v := range rank {
			if v {
				sum++
			}
		}
	}
	return sum
}
