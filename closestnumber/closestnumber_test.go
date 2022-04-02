package main

import (
	"testing"

	"github.com/act28/hackerrank.git/capturer"
	"github.com/stretchr/testify/assert"
)

func TestClosestNumber(t *testing.T) {
	testCases := []struct {
		in  []int32
		exp []string
	}{
		{
			in:  []int32{4, 2, 1, 3},
			exp: []string{"1 2", "2 3", "3 4"},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			output := capturer.CaptureOutput(func() {
				closestNumbers(tc.in)
			})
			assert.Equal(t, tc.exp, output)
		})
	}
}
