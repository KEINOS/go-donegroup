package donegroup

import (
	"math/bits"

	"github.com/pkg/errors"
)

const NumFlagPerUnitMax = 64 // Max number of flags in a group/unit
const MaskAll = ^uint64(0)   // Mask to set all flags to done

// ----------------------------------------------------------------------------
//  Type: DoneGroup
// ----------------------------------------------------------------------------

// DoneGroup is a group of flags that can be set to done/true or undone/false.
type DoneGroup struct {
	flagUnits []uint64 // List of flags. Each flag unit holds 64
	numFlags  int      // Number of flags to manage
	numUnit   int      // Number of flag units
}

// ----------------------------------------------------------------------------
//  Constructor
// ----------------------------------------------------------------------------

func New(numFlags int) (*DoneGroup, error) {
	if numFlags <= 0 {
		return nil, errors.New("number of flags to manage must be greater than 0")
	}

	numUnit := (numFlags / NumFlagPerUnitMax) + 1

	return &DoneGroup{
		flagUnits: make([]uint64, numUnit),
		numFlags:  numFlags,
		numUnit:   numUnit,
	}, nil
}

// ----------------------------------------------------------------------------
//  Methods
// ----------------------------------------------------------------------------

// done is the actual function of Done(). It sets the flag at the given index
// to done/true.
func (dg *DoneGroup) done(index int) error {
	unitIndex, posIndex, err := dg.getPos(index)
	if err != nil {
		return errors.Wrap(err, "failed to set flag to done")
	}

	dg.flagUnits[unitIndex] |= GetMask(posIndex)

	return nil
}

// Done sets the flag at the given index to done/true.
func (dg *DoneGroup) Done(index int) bool {
	if err := dg.done(index); err != nil {
		return false
	}

	return true
}

// DoneAll sets all flags to done/true.
func (dg *DoneGroup) DoneAll() {
	for i := 0; i < dg.numUnit; i++ {
		dg.flagUnits[i] = MaskAll
	}
}

// getPos returns the actual position of the flag in the flag unit.
func (dg DoneGroup) getPos(index int) (unitIndex int, posIndex int, err error) {
	if index <= 0 || index > dg.numFlags {
		return 0, 0, errors.Errorf("index out of range: %d-%d", 1, dg.numFlags)
	}

	if index < NumFlagPerUnitMax {
		return 0, index, nil
	}

	index -= 1 // 0-based index

	// Number of the flag unit
	unitIndex = index / NumFlagPerUnitMax
	// Index in the unit. It will iterate between 1-NumFlagPerUnitMax
	posIndex = (index % NumFlagPerUnitMax) + 1

	return unitIndex, posIndex, nil
}

// IsDone returns true if the flag at the given index is done/true.
func (dg *DoneGroup) IsDone(index int) bool {
	unitIndex, posIndex, err := dg.getPos(index)
	if err != nil {
		return false
	}

	return dg.flagUnits[unitIndex]&GetMask(posIndex) != 0
}

// IsDoneAll returns true if all flags are done/true.
func (dg *DoneGroup) IsDoneAll() bool {
	total := 0

	for i := 0; i < dg.numUnit; i++ {
		total += bits.OnesCount64(dg.flagUnits[i])
	}

	return total >= dg.numFlags
}

// MustDone is the same as Done() but panics if an error occurs.
func (dg *DoneGroup) MustDone(index int) {
	if err := dg.done(index); err != nil {
		panic(err)
	}
}

// MustUndone is the same as Undone() but panics if an error occurs.
func (dg *DoneGroup) MustUndone(index int) {
	if err := dg.undone(index); err != nil {
		panic(err)
	}
}

// undone is the actual function of Undone(). It sets the flag at the given
// index to undone/false.
func (dg *DoneGroup) undone(index int) error {
	unitIndex, posIndex, err := dg.getPos(index)
	if err != nil {
		return errors.Wrap(err, "failed to set flag to undone")
	}

	dg.flagUnits[unitIndex] ^= GetMask(posIndex)

	return nil
}

// Undone sets the flag at the given index to undone/false.
func (dg *DoneGroup) Undone(index int) bool {
	if err := dg.undone(index); err != nil {
		return false
	}

	return true
}

// UndoneAll sets all flags to undone/false.
func (dg *DoneGroup) UndoneAll() {
	for i := 0; i < dg.numUnit; i++ {
		dg.flagUnits[i] = 0
	}
}
