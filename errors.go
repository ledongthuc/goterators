package goterators

import "fmt"

var (
	// ErrNotFound is used in functions that need to find any items in the list but are not found.
	ErrNotFound = fmt.Errorf("not found")
)
