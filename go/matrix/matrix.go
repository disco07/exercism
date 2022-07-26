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
	if len(*m) == 0 {
		return [][]int{}
	}
	arr := make([][]int, 0)
	for j := 0; j < len((*m)[0]); j++ {
		r := make([]int, 0)
		for i := 0; i < len(*m); i++ {
			r = append(r, (*m)[i][j])
		}
		arr = append(arr, r)
	}
	return arr
}

func (m *Matrix) Rows() [][]int {
	arr := make([][]int, len(*m))
	for i := 0; i < len(*m); i++ {
		r := make([]int, len((*m)[0]))
		copy(r, (*m)[i])
		arr[i] = r
	}
	return arr
}

func (m *Matrix) Set(row, col, val int) bool {
	if len(*m) == 0 || row >= len(*m) || col >= len((*m)[0]) || row < 0 || col < 0 {
		return false
	}
	(*m)[row][col] = val
	return true
}
