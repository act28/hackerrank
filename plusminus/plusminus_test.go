package main

import (
	"testing"

	"github.com/act28/hackerrank.git/capturer"
	"github.com/stretchr/testify/assert"
)

func TestPlusMinus(t *testing.T) {
	testCases := []struct {
		in  []int32
		exp []string
	}{
		{
			in:  []int32{-4, 3, -9, 0, 4, 1},
			exp: []string{"0.500000", "0.333333", "0.166667"},
		},
		{
			in:  []int32{1, 2, 3, -1, -2, -3, 0, 0},
			exp: []string{"0.375000", "0.375000", "0.250000"},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			output := capturer.CaptureOutput(func() {
				plusMinus(tc.in)
			})
			assert.Equal(t, tc.exp, output)
		})
	}
}
