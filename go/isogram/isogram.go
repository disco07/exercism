package isogram

import "strings"

func IsIsogram(word string) bool {
	var count int
	for _, w := range []rune(word) {
		count = 0
		for _, w2 := range []rune(word) {
			if strings.ToLower(string(w)) == strings.ToLower(string(w2)) && string(w) != "-" && string(w) != " " {
				count++
			}
			if count > 1 {
				return false
			}
		}
	}
	return true
}
