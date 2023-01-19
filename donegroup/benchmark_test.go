package donegroup

import (
	"testing"

	"github.com/Code-Hex/dd"
	"github.com/stretchr/testify/require"
)

func BenchmarkDone(b *testing.B) {
	const numFlags = 100

	dg, err := New(numFlags)
	require.NoError(b, err, "failed to initialize DoneGroup")

	b.ResetTimer()
	for ite := 0; ite < b.N; ite++ {
		// Loop 1-100
		for index := 1; index <= numFlags; index++ {
			if !dg.Done(index) {
				b.Fatalf("failed to set flag to done: index: %d, structure: %s",
					index, dd.Dump(dg))
			}
		}
	}
}
