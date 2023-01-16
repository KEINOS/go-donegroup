package donegroup_test

import (
	"fmt"
	"log"

	"github.com/KEINOS/go-donegroup/donegroup"
)

func Example() {
	// ----------------------------------------------------------------------------
	//  Instantiate a DoneGroup (Constructor)
	// ----------------------------------------------------------------------------

	// Prepare a DoneGroup with 3 flags (for index range: 1-3)
	doneGroup, err := donegroup.New(3)
	if err != nil {
		log.Fatal(err)
	}

	// ----------------------------------------------------------------------------
	//  Set the flag as done (flag up, set to true)
	// ----------------------------------------------------------------------------

	// Done will flag up the first flag (set index 1 to true).
	// It returns true on success and false on error (out of range for example).
	if doneGroup.Done(1) {
		fmt.Println("Flag 1 is now up")
	} else {
		fmt.Println("failed to flag up the flag")
	}

	// ----------------------------------------------------------------------------
	//  Check the flag status by index
	// ----------------------------------------------------------------------------

	// Check the status of the flag of index 1
	if doneGroup.IsDone(1) {
		fmt.Println("Flag 1 is up") // <-- expected
	} else {
		fmt.Println("Flag 1 is down")
	}

	// Check the status of the flag of index 100 which is out of range
	if doneGroup.IsDone(100) {
		fmt.Println("Flag 100 is up")
	} else {
		fmt.Println("Flag 100 is down") // <-- expected
	}

	// ----------------------------------------------------------------------------
	//  Un-set the flag as undone (flag down, set to false)
	// ----------------------------------------------------------------------------

	// Un-done/flag down the first flag (set index 0 to false)
	if doneGroup.Undone(1) {
		fmt.Println("Flag 1 is now down")
	}

	// Check the status of the flag of index 1
	if doneGroup.IsDone(1) {
		fmt.Println("Flag 1 is up")
	} else {
		fmt.Println("Flag 1 is down") // <-- expected
	}

	if doneGroup.Undone(100) {
		fmt.Println("Flag 100 is now down")
	} else {
		fmt.Println("failed to flag down the flag") // <-- expected
	}

	// ----------------------------------------------------------------------------
	//  Check and manage all of the flag status
	// ----------------------------------------------------------------------------

	// Check if all the flags are up
	if doneGroup.IsDoneAll() {
		fmt.Println("All flags are up")
	} else {
		fmt.Println("Not all flags are up") // <-- expected
	}

	// Set all the flags up (set all the flags to true)
	doneGroup.DoneAll()

	// Check if all the flags are up
	if doneGroup.IsDoneAll() {
		fmt.Println("All flags are up") // <-- expected
	} else {
		fmt.Println("Not all flags are up")
	}

	// Set all the flags down (set all the flags to false)
	doneGroup.UndoneAll()

	// Check if all the flags are up
	if doneGroup.IsDoneAll() {
		fmt.Println("All flags are up")
	} else {
		fmt.Println("Not all flags are up") // <-- expected
	}

	// ----------------------------------------------------------------------------
	//  Out of range
	// ----------------------------------------------------------------------------

	// Setting a flag out of range will return false
	if doneGroup.Done(100) {
		fmt.Println("Flag 100 is now up")
	} else {
		fmt.Println("Flag 100 is out of range") // <-- expected
	}

	// ----------------------------------------------------------------------------
	//  Must functions
	// ----------------------------------------------------------------------------

	// MustDone will panic on error. 1 is the index in-range of 3 so it will not panic.
	doneGroup.MustDone(1)

	// MustUndone will panic on error. 1 and 3 is the index in-range of 3 so it will not panic.
	doneGroup.MustUndone(1)
	doneGroup.MustUndone(3)

	// Output:
	// Flag 1 is now up
	// Flag 1 is up
	// Flag 100 is down
	// Flag 1 is now down
	// Flag 1 is down
	// failed to flag down the flag
	// Not all flags are up
	// All flags are up
	// Not all flags are up
	// Flag 100 is out of range
}
