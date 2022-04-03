package main

import (
	"testing"

	"github.com/act28/hackerrank/problemsolving/capturer"
	"github.com/stretchr/testify/assert"
)

func TestCountApplesAndOranges(t *testing.T) {
	testCases := []struct {
		s       int32
		t       int32
		a       int32
		b       int32
		apples  []int32
		oranges []int32
		exp     []string
	}{
		{
			s:       7,
			t:       10,
			a:       4,
			b:       12,
			apples:  []int32{2, 3, -4},
			oranges: []int32{3, -2, -4},
			exp:     []string{"1", "2"},
		},
		{
			s:       7,
			t:       11,
			a:       5,
			b:       15,
			apples:  []int32{-2, 2, 1},
			oranges: []int32{5, -6},
			exp:     []string{"1", "1"},
		},
	}

	for _, tc := range testCases {
		t.Run("count apples and oranges", func(t *testing.T) {
			output := capturer.CaptureOutput(func() {
				countApplesAndOranges(tc.s, tc.t, tc.a, tc.b, tc.apples, tc.oranges)
			})
			assert.Equal(t, tc.exp, output)
		})
	}
}
