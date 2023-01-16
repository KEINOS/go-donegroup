package donegroup

import (
	"testing"

	"github.com/Code-Hex/dd"
	"github.com/stretchr/testify/require"
)

func BenchmarkDone(b *testing.B) {
	numFlags := NumFlagPerUnitMax * b.N // 64*N

	dg, err := New(numFlags)
	require.NoError(b, err, "failed to initialize DoneGroup")

	const index = 1

	b.ResetTimer()
	for ite := 0; ite < b.N; ite++ {
		if !dg.Done(index) {
			b.Fatalf("failed to set flag to done: index: %d, structure: %s",
				index, dd.Dump(dg))
		}
	}
}
