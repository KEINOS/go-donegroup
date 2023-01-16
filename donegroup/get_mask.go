package donegroup

// GetMask returns a mask with a single bit set at the given position.
func GetMask(posIndex int) uint64 {
	if posIndex <= 0 || posIndex > NumFlagPerUnitMax {
		return 0
	}

	return 1 << (posIndex - 1)
}
