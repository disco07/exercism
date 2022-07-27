package luhn

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	if num, isDoubling := doubling(id); isDoubling == true {
		return sum(num)
	}
	return false
}

func doubling(id string) ([]int, bool) {
	id = strings.TrimSpace(id)
	var i2 int
	var num []int

	for _, r := range []rune(id) {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, false
		}
		num = append(num, n)
	}
	fmt.Println(num)

	if len(num) <= 1 {
		return nil, false
	}

	num = sliceReverse(num)
	for i, arr := range num {
		if (i+1)%2 == 0 {
			i2 = arr * 2
			if i2 > 9 {
				i2 -= 9
			}
			num[i] = i2
		}
	}
	fmt.Println(sliceReverse(num), num)
	return sliceReverse(num), true
}

func sum(num []int) bool {
	var count int
	for _, i := range num {
		count += i
	}
	return count%10 == 0
}

func sliceReverse(slc []int) []int {
	var revSlc []int
	for i := range slc {
		revSlc = append(revSlc, slc[len(slc)-1-i])
	}
	return revSlc
}
