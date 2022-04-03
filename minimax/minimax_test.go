package main

import (
	"testing"

	"github.com/act28/hackerrank/capturer"
	"github.com/stretchr/testify/assert"
)

func TestMiniMax(t *testing.T) {
	testCases := []struct {
		in  []int32
		exp []string
	}{
		{
			in:  []int32{1, 3, 5, 7, 9},
			exp: []string{"16 24"},
		},
		{
			in:  []int32{1, 2, 3, 4, 5},
			exp: []string{"10 14"},
		},
		{
			in:  []int32{1, 2, 3, 4, 5, 6},
			exp: []string{"15 20"},
		},
		{
			in:  []int32{7, 69, 2, 221, 8974},
			exp: []string{"299 9271"},
		},
	}

	for _, tc := range testCases {
		t.Run("minimax", func(t *testing.T) {
			output := capturer.CaptureOutput(func() {
				miniMaxSum(tc.in)
			})
			assert.Equal(t, tc.exp, output)
		})
	}
}
