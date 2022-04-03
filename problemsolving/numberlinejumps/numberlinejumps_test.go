package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberLineJumps(t *testing.T) {
	testCases := []struct {
		x1  int32
		v1  int32
		x2  int32
		v2  int32
		exp string
	}{
		{
			x1:  2,
			v1:  1,
			x2:  1,
			v2:  2,
			exp: "YES",
		},
		{
			x1:  2,
			v1:  1,
			x2:  3,
			v2:  1,
			exp: "NO",
		},
		{
			x1:  0,
			v1:  3,
			x2:  4,
			v2:  2,
			exp: "YES",
		},
		{
			x1:  0,
			v1:  2,
			x2:  5,
			v2:  3,
			exp: "NO",
		},
		{
			x1:  1928,
			v1:  4306,
			x2:  5763,
			v2:  4301,
			exp: "YES",
		},
		{
			x1:  7271,
			v1:  2211,
			x2:  7915,
			v2:  2050,
			exp: "YES",
		},
	}

	for _, tc := range testCases {
		t.Run("number line jumps", func(t *testing.T) {
			assert.Equal(t, tc.exp, kangaroo(tc.x1, tc.v1, tc.x2, tc.v2))
		})
	}
}
