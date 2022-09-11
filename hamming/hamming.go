package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var count = 0
	if len(a) != len(b) {
		return count, errors.New("length of sequences are not equal")
	}
	for i, r := range []rune(a) {
		if r != []rune(b)[i] {
			count++
		}
	}
	return count, nil
}
