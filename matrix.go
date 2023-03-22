package main

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func CreateFromCSV(records [][]string) (Matrix, error) {
	rowsN := len(records)
	// preallocate matrix.
	inverted := make(Matrix, rowsN)
	for i := range records {
		inverted[i] = make([]int, len(records[i]))
	}
	for i := range records {
		columnN := len(records[i])
		if rowsN != columnN {
			return nil, errors.New("matrix should be square")
		}
		for j := range records[i] {
			digit, err := strconv.Atoi(records[i][j])
			if err != nil {
				return nil, errors.New("all matrix elems should be numbers")
			}
			inverted[i][j] = digit
		}
	}
	return inverted, nil
}

func (m Matrix) Echo() string {
	rows := make([]string, 0, len(m))
	for i := range m {
		rowValues := make([]string, 0, len(m[i]))
		for j := range m[i] {
			rowValues = append(rowValues, strconv.Itoa(m[i][j]))
		}
		rows = append(rows, strings.Join(rowValues, ","))
	}
	return strings.Join(rows, "\n")
}

func (m Matrix) Mul() int {
	result := 1
	for i := range m {
		for j := range m[i] {
			result *= m[i][j]
		}
	}
	return result
}

func (m Matrix) Sum() int {
	result := 0
	for i := range m {
		for j := range m[i] {
			result += m[i][j]
		}
	}
	return result
}

func (m Matrix) Flatten() string {
	values := make([]string, 0, len(m)^2)
	for i := range m {
		for j := range m[i] {
			values = append(values, strconv.Itoa(m[i][j]))
		}
	}
	return strings.Join(values, ",")
}

func (m Matrix) Invert() Matrix {
	// preallocate matrix.
	inverted := make(Matrix, len(m))
	for i := range m {
		inverted[i] = make([]int, len(m[i]))
	}
	for i := range m {
		for j := range m[i] {
			inverted[j][i] = m[i][j]
		}
	}
	return inverted
}
