package inhouse

import (
	"sort"
)

// Code represents a Go code with some information for post processing.
type Code struct {
	Filepath string
	Function string
	Line     int
}

// Sort slice of code by filename.
func sortCode(given []*Code) []*Code {
	sort.Slice(given[:], func(i, j int) bool {
		return given[i].Filepath < given[j].Filepath
	})

	return given
}
