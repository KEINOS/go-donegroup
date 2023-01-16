package donegroup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetMask(t *testing.T) {
	t.Parallel()

	for numTest, data := range []struct {
		input      int
		expectMask uint64
	}{
		{input: 0, expectMask: 0},
		{input: 65, expectMask: 0},
		{input: 1, expectMask: 0b1},
		{input: 10, expectMask: 0b1000000000},
		{input: 32, expectMask: 0b10000000000000000000000000000000},
		{
			input:      64,
			expectMask: 0b1000000000000000000000000000000000000000000000000000000000000000,
		},
	} {
		expectMask := data.expectMask
		actualMask := GetMask(data.input)

		require.Equal(t, expectMask, actualMask, "test #%d: input: %d", numTest, data.input)
	}
}
