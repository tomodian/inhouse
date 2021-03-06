package inhouse

import (
	"encoding/json"
)

// ContainsOutput represents a matching result of code.
type ContainsOutput struct {
	Contained bool
	Code      *Code
}

// HasCode internally checks code field.
func (c ContainsOutput) HasCode() bool {
	return c.Code != nil
}

// Check represents a checker result.
type Check struct {
	Contained bool    `json:"contained"`
	Matches   []*Code `json:"matches"`
	Misses    []*Code `json:"misses"`

	// Just for testing.
	Tester interface{} `json:"tester,omitempty"`
}

// NewCheck with initialized slices.
func NewCheck() *Check {
	return &Check{
		Matches: []*Code{},
		Misses:  []*Code{},
	}
}

// Sort internal codes.
func (c *Check) Sort() {
	if len(c.Matches) > 0 {
		c.Matches = sortCode(c.Matches)
	}

	if len(c.Misses) > 0 {
		c.Misses = sortCode(c.Misses)
	}
}

// Combine all codes.
func (c *Check) Combine() []*Code {
	out := []*Code{}

	out = append(out, c.Matches...)
	out = append(out, c.Misses...)

	return sortCode(out)
}

// ToJSON for CLI output.
func (c Check) ToJSON() (string, error) {

	b, err := json.Marshal(c)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Contains returns true when source files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will search for source and test files.
func Contains(dir, function string, recursive bool) (*Check, error) {
	files, err := Sources(dir, recursive)

	if err != nil {
		return nil, err
	}

	tests, err := Tests(dir, recursive)

	if err != nil {
		return nil, err
	}

	merged := []string{}
	merged = append(merged, files...)
	merged = append(merged, tests...)

	return parseContains(function, merged)
}

// ContainsPWD returns true when source files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will search for source and test files.
func ContainsPWD(function string, recursive bool) (*Check, error) {
	pwd, err := PWD()

	if err != nil {
		return nil, err
	}

	files, err := Sources(pwd, recursive)

	if err != nil {
		return nil, err
	}

	tests, err := Tests(pwd, recursive)

	if err != nil {
		return nil, err
	}

	merged := []string{}
	merged = append(merged, files...)
	merged = append(merged, tests...)

	return parseContains(function, merged)
}

// SourcesContains returns true when source files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will not check for `*_test.go` files.
func SourcesContains(dir, function string, recursive bool) (*Check, error) {
	files, err := Sources(dir, recursive)

	if err != nil {
		return nil, err
	}

	return parseContains(function, files)
}

// SourcesContainsPWD is a wrapper of SourcesContains to resolve directory as `pwd`.
// Intended to be used inside test codes.
func SourcesContainsPWD(function string, recursive bool) (*Check, error) {
	pwd, err := PWD()

	if err != nil {
		return nil, err
	}

	return SourcesContains(pwd, function, recursive)
}

// TestsContains returns true when test files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will check for `*_test.go` files only.
func TestsContains(dir, function string, recursive bool) (*Check, error) {
	files, err := Tests(dir, recursive)

	if err != nil {
		return nil, err
	}

	return parseContains(function, files)
}

// TestsContainsPWD is a wrapper of TestsContains to resolve directory as `pwd`.
// Intended to be used inside test codes.
func TestsContainsPWD(function string, recursive bool) (*Check, error) {
	pwd, err := PWD()

	if err != nil {
		return nil, err
	}

	return TestsContains(pwd, function, recursive)
}

// matchFunction returns true when Go source file contains the specified Go function.
// `function` parameter is case sensitive, since it checks for both exported and private functions.
// `error` will be returned in illogical situation, such as passing Non-Go files (e.g. Markdown) or file is missing.
// Therefore you should always handle the error at first, then proceed to bool checks.
func matchFunction(src, function string) (*ContainsOutput, error) {

	codes, err := Functions(src)

	if err != nil {
		return nil, err
	}

	out := &ContainsOutput{}

	for _, c := range codes {
		out.Code = &c

		if c.Function == function {
			out.Contained = true
			continue
		}
	}

	return out, nil
}

// parseContains return a generic check result.
func parseContains(function string, files []string) (*Check, error) {
	out := NewCheck()

	for _, f := range files {
		c, err := matchFunction(f, function)

		if err != nil {
			return nil, err
		}

		if !c.HasCode() {
			continue
		}

		if c.Contained {
			out.Contained = true
			out.Matches = append(out.Matches, c.Code)

			continue
		}

		out.Misses = append(out.Misses, c.Code)
	}

	out.Sort()

	return out, nil
}
