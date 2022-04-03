package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressingArray(t *testing.T) {
	testCases := []struct {
		in  []int32
		max int32
		exp int32
	}{
		{
			in:  []int32{3, 5, 1, 4, 5},
			max: 15,
			exp: 3,
		},
		{
			in:  []int32{3, 5, 1, 4, 5},
			max: 10,
			exp: 5,
		},
		{
			in:  []int32{3, 5, 1, 4, 5},
			max: 20,
			exp: 3,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 1},
			max: 60,
			exp: 3,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 1},
			max: 100,
			exp: 3,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 4},
			max: 500,
			exp: 2,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 4},
			max: 600,
			exp: 2,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 4},
			max: 1500,
			exp: 1,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 4},
			max: 25000,
			exp: 1,
		},
		{
			in:  []int32{3, 5, 1, 4, 5, 4, 3},
			max: 125000,
			exp: 1,
		},
	}

	for _, tc := range testCases {
		t.Run("compressing array", func(t *testing.T) {
			output := getMinLength(tc.in, tc.max)
			assert.Equal(t, tc.exp, output, tc.max)
		})
	}
}
