package donegroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
//  DoneGroup.getPos()
// ----------------------------------------------------------------------------

// DataSpecGetPos is the data provider to check the required spec for getPos
// method.
var DataSpecGetPos = []struct {
	indexIn    int
	unitIndex  int
	posIndex   int
	requireErr bool
}{
	{indexIn: 0, unitIndex: 0, posIndex: 0, requireErr: true},
	{indexIn: 1, unitIndex: 0, posIndex: 1, requireErr: false},
	{indexIn: 63, unitIndex: 0, posIndex: 63, requireErr: false},
	{indexIn: 64, unitIndex: 0, posIndex: 64, requireErr: false},
	{indexIn: 65, unitIndex: 1, posIndex: 1, requireErr: false},
	{indexIn: 128, unitIndex: 1, posIndex: 64, requireErr: false},
	{indexIn: 129, unitIndex: 2, posIndex: 1, requireErr: false},
	{indexIn: 130, unitIndex: 0, posIndex: 0, requireErr: true},
}

func TestDoneGroup_getPos(t *testing.T) {
	t.Parallel()

	dg, err := New(NumFlagPerUnitMax*2 + 1)
	require.NoError(t, err, "failed to initialize DoneGroup")

	for numTest, testData := range DataSpecGetPos {
		actualUnitIndex, actualPosIndex, actualErr := dg.getPos(testData.indexIn)

		if testData.requireErr {
			require.Error(t, actualErr, "flag #%d should error", numTest+1)
		} else {
			require.NoError(t, actualErr, "flag #%d should not error", numTest+1)
		}

		expectUnitIndex := testData.unitIndex
		expectPosIndex := testData.posIndex

		assert.Equal(t, expectUnitIndex, actualUnitIndex,
			"test #%d unit index failed: input index: %d", numTest+1, testData.indexIn)
		assert.Equal(t, expectPosIndex, actualPosIndex,
			"test #%d pos index failed: input index: %d", numTest+1, testData.indexIn)
	}
}

// ----------------------------------------------------------------------------
//  DoneGroup.MustDone()
// ----------------------------------------------------------------------------

func TestDoneGroup_MustDone(t *testing.T) {
	t.Parallel()

	dg, err := New(3)
	require.NoError(t, err, "failed to initialize DoneGroup during test")

	require.Panics(t, func() {
		dg.MustDone(100)
	}, "MustDone should panic on error")
}

// ----------------------------------------------------------------------------
//  DoneGroup.MustUndone()
// ----------------------------------------------------------------------------

func TestDoneGroup_MustUndone(t *testing.T) {
	t.Parallel()

	dg, err := New(3)
	require.NoError(t, err, "failed to initialize DoneGroup during test")

	require.Panics(t, func() {
		dg.MustUndone(100)
	}, "MustDone should panic on error")
}

// ----------------------------------------------------------------------------
//  DoneGroup.undone()
// ----------------------------------------------------------------------------

// This test detects the XOR bug in undone() method. Calling more than once
// consecutively with the same index should not turn true.
func TestDoneGroup_undone_twice(t *testing.T) {
	t.Parallel()

	dg, err := New(64) // flag range 1-64
	require.NoError(t, err, "failed to initialize DoneGroup during test")

	require.False(t, dg.IsDone(3), "it should false if not set yet")

	require.NoError(t, dg.undone(3), "it should not error if index is in range(1-64)")
	require.False(t, dg.IsDone(3), "it should be false if undone success with the same index")

	// Call undone() again
	require.NoError(t, dg.undone(3), "it should not error if index is in range(1-64)")
	require.False(t, dg.IsDone(3), "it should not turn true if undone success with the same index")
}

// ----------------------------------------------------------------------------
//  New()
// ----------------------------------------------------------------------------

func TestNew_fail_instantiate(t *testing.T) {
	t.Parallel()

	for numTest, testData := range []struct {
		numFlags int
	}{
		{numFlags: 0},
		{numFlags: -1},
		{numFlags: -100},
	} {
		dg, err := New(testData.numFlags)

		require.Error(t, err, "Test %d: expected error", numTest)
		assert.Nil(t, dg, "Test %d: expected nil", numTest)
	}
}
