package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// matrix is a data type that represents a 2D matrix.
type matrix [][]int

// New constructs a matrix based on the input string.
func New(s string) (*matrix, error) {
	var size int
	lines := strings.Split(s, "\n")
	m := make(matrix, len(lines))

	for i, line := range lines {
		vals := strings.Split(strings.TrimSpace(line), " ")
		if i == 0 {
			size = len(vals)
		} else if size != len(vals) {
			return nil, errors.New("wrong number of values in line")
		}
		m[i] = make([]int, len(vals))
		for j, val := range vals {
			v, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			m[i][j] = v
		}
	}
	return &m, nil
}

// Rows returns a row representation of the matrix.
func (m *matrix) Rows() [][]int {
	d := make([][]int, len(*m))
	for i := range *m {
		d[i] = make([]int, len((*m)[i]))
		copy(d[i], (*m)[i])
	}
	return d
}

// Cols returns a column representation of the matrix.
func (m *matrix) Cols() [][]int {
	t := make([][]int, len((*m)[0]))
	for x := range t {
		t[x] = make([]int, len(*m))
	}
	for i, row := range *m {
		for j, val := range row {
			t[j][i] = val
		}
	}
	return t
}

// Set sets the value of an element in the matrix.
func (m *matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) {
		return false
	}
	if col < 0 || col >= len((*m)[0]) {
		return false
	}
	(*m)[row][col] = val
	return true
}
