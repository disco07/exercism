package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	m := make(Matrix, 0)
	for i, row := range rows {
		r := strings.Fields(row)
		if i != 0 && len(r) != len(m[i-1]) {
			return nil, errors.New("non even matrix")
		}
		arr := make([]int, 0)
		for _, e := range r {
			v, err := strconv.Atoi(e)
			if err != nil {
				return nil, errors.New("invalid integer")
			}
			arr = append(arr, v)
		}
		m = append(m, arr)
	}
	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	panic("Please implement the Cols function")
}

func (m *Matrix) Rows() [][]int {
	panic("Please implement the Rows function")
}

func (m *Matrix) Set(row, col, val int) bool {
	panic("Please implement the Set function")
}
