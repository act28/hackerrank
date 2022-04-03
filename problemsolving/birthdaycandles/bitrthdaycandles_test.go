package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	testCases := []struct {
		in  []int32
		exp int32
	}{
		{
			in:  []int32{4, 4, 2, 1},
			exp: 2,
		},
		{
			in:  []int32{3, 2, 1, 3},
			exp: 2,
		},
	}

	for _, tc := range testCases {
		t.Run("birthday cake candles", func(t *testing.T) {
			assert.Equal(t, tc.exp, birthdayCakeCandles(tc.in))
		})
	}
}
