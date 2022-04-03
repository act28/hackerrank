package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimneConversion(t *testing.T) {
	testCases := []struct {
		in  string
		exp string
	}{
		{
			in:  "12:01:00PM",
			exp: "12:01:00",
		},
		{
			in:  "12:01:00AM",
			exp: "00:01:00",
		},
		{
			in:  "07:05:45PM",
			exp: "19:05:45",
		},
		{
			in:  "09:55:45PM",
			exp: "21:55:45",
		},
		{
			in:  "11:59:59PM",
			exp: "23:59:59",
		},
	}

	for _, tc := range testCases {
		t.Run("time conversion", func(t *testing.T) {
			assert.Equal(t, tc.exp, timeConversion(tc.in))
		})
	}
}
