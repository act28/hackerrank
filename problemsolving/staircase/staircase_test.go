package main

import (
	"testing"

	"github.com/act28/hackerrank/capturer"
	"github.com/stretchr/testify/assert"
)

func TestStaircase(t *testing.T) {
	testCases := []struct {
		in  int32
		exp []string
	}{
		{
			in:  4,
			exp: []string{"   #", "  ##", " ###", "####"},
		},
		{
			in:  5,
			exp: []string{"    #", "   ##", "  ###", " ####", "#####"},
		},
		{
			in:  10,
			exp: []string{"         #", "        ##", "       ###", "      ####", "     #####", "    ######", "   #######", "  ########", " #########", "##########"},
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			output := capturer.CaptureOutput(func() {
				staircase(tc.in)
			})
			assert.Equal(t, tc.exp, output)
		})
	}
}
