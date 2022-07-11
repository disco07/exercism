package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if fmt.Sprintf("%T", fnb) == "sorting.FancyNumber" {
		conv, err := strconv.ParseInt(fnb.Value(), 0, 64)
		if err != nil {
			return 0
		}
		return int(conv)
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	if fmt.Sprintf("%T", i) == "string" {
		return "Return to sender"
	}
	switch i.(type) {
	case FancyNumber:
		f, _ := i.(FancyNumber)
		return DescribeFancyNumberBox(f)
	case int, float64:
		f, ok := i.(float64)
		if !ok {
			conv, _ := i.(int)
			f = float64(conv)
		}
		return fmt.Sprintf("This is the number %.1f", f)
	case string:
		return "Return to sender"
	default:
		f, _ := i.(FancyNumber)
		return DescribeFancyNumberBox(f)
	}
}
