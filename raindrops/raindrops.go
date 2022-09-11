package raindrops

import "fmt"

func Convert(number int) string {
	var s = ""
	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = fmt.Sprintf("%v", number)
	}
	return s
}
