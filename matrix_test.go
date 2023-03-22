package main

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func matrixFactory() Matrix {
	return Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
}

func TestCreateFromCSV(t *testing.T) {
	tests := []struct {
		name        string
		input       [][]string
		expected    Matrix
		expectedErr error
	}{
		{
			name: "fail: matrix is not square",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6", "7"},
				{"8", "9", "10"},
			},
			expected:    nil,
			expectedErr: errors.New("matrix should be square"),
		},
		{
			name: "fail: elem of matrix is not a number",
			input: [][]string{
				{"1", "bla", "3"},
				{"4", "5", "bloom"},
				{"8", "9", "10"},
			},
			expected:    nil,
			expectedErr: errors.New("all matrix elems should be numbers"),
		},
		{
			name: "success",
			input: [][]string{
				{"1", "2", "3"},
				{"4", "5", "6"},
				{"7", "8", "9"},
			},
			expected: Matrix{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expectedErr: nil,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := CreateFromCSV(tc.input)
			require.Equal(t, tc.expected, actual)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestMatrix_Echo(t *testing.T) {
	m := matrixFactory()
	expected := `1,2,3
4,5,6
7,8,9`
	actual := m.Echo()
	require.Equal(t, expected, actual)
}

func TestMatrix_Mul(t *testing.T) {
	m := matrixFactory()
	expected := 362880
	actual := m.Mul()
	require.Equal(t, expected, actual)
}

func TestMatrix_Sum(t *testing.T) {
	m := matrixFactory()
	expected := 45
	actual := m.Sum()
	require.Equal(t, expected, actual)
}

func TestMatrix_Flatten(t *testing.T) {
	m := matrixFactory()
	expected := "1,2,3,4,5,6,7,8,9"
	actual := m.Flatten()
	require.Equal(t, expected, actual)
}

func TestMatrix_Invert(t *testing.T) {
	m := matrixFactory()
	expected := Matrix{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	actual := m.Invert()
	require.Equal(t, expected, actual)
}
