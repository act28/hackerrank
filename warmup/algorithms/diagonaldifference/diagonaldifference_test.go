package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDifference(t *testing.T) {
	testCases := []struct {
		in  [][]int32
		exp int32
	}{
		{
			in:  [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}},
			exp: 15,
		},
		{
			in:  [][]int32{{11, 2, 4, 6}, {4, 5, 6, -2}, {10, 8, -12, 7}, {-3, 5, 7, 9}},
			exp: 4,
		},
	}

	for _, tc := range testCases {
		t.Run("diagonal difference", func(t *testing.T) {
			output := diagonalDifference(tc.in)
			assert.Equal(t, tc.exp, output)
		})
	}
}
