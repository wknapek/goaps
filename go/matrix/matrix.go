package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for idx, row := range m {
		rows[idx] = make([]int, len(row))
		copy(rows[idx], row)
	}
	return rows
}

func New(m string) (Matrix, error) {
	Rows := strings.Split(m, "\n")

	matrix := make([][]int, len(Rows))
	for ridx, row := range Rows {
		rowelem := strings.Split(strings.TrimSpace(row), " ")
		if ridx != 0 && len(matrix[ridx-1]) != len(rowelem) {
			return nil, errors.New("wrong row lengt")
		}
		matrix[ridx] = make([]int, len(rowelem))
		for cidx, col := range rowelem {
			converted, err := strconv.Atoi(col)
			if err != nil {
				return nil, err
			}
			matrix[ridx][cidx] = converted
		}

	}
	return matrix, nil
}

func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for _, row := range m {
		for idx, col := range row {
			cols[idx] = append(cols[idx], col)
		}
	}
	return cols
}

func (m Matrix) Set(row int, col int, val int) bool {
	if row < 0 || row > len(m)-1 || col < 0 || col > len(m[0])-1 {
		return false
	}
	m[row][col] = val
	return true
}
