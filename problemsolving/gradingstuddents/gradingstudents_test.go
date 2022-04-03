package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGradingStudents(t *testing.T) {
	testCases := []struct {
		in  []int32
		exp []int32
	}{
		{
			in:  []int32{73, 67, 38, 33},
			exp: []int32{75, 67, 40, 33},
		},
	}

	for _, tc := range testCases {
		t.Run("grading students", func(t *testing.T) {
			assert.Equal(t, tc.exp, gradingStudents(tc.in))
		})
	}
}
