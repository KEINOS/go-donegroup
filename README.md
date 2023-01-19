# go-donegroup

`go-donegroup` is a Go package to manage bit flags.

Which is somewhat similar to `sync.WaitGroup` but with an index.

## Usage

```go
go get "github.com/KEINOS/go-donegroup"
```

```go
import "github.com/KEINOS/go-donegroup/donegroup"
```

### Example

```go
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
```

## Benchstat

```text
goos: darwin
goarch: amd64
pkg: github.com/KEINOS/go-donegroup/donegroup
cpu: Intel(R) Core(TM) i5-5257U CPU @ 2.70GHz
name             time/op         alloc/op       allocs/op
BenchmarkDone-4  7.48ns ± 0%     0.00B ± 0%     0 allocs/op ± 0%
```
