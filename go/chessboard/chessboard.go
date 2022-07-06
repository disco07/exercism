package chessboard

// Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []struct {
	bool
}

// Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard struct {
	map[string]Rank
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	panic("Please implement CountInRank()")
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	panic("Please implement CountInFile()")
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
