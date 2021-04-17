package inhouse

// ContainsOutput represents a matching result of code.
type ContainsOutput struct {
	Contained bool
	Code      *Code
}

// HasCode internally checks code field.
func (c ContainsOutput) HasCode() bool {
	return c.Code != nil
}

// Contains returns true when Go source file contains the specified Go function.
// `function` parameter is case sensitive, since it checks for both exported and private functions.
// `error` will be returned in illogical situation, such as passing Non-Go files (e.g. Markdown) or file is missing.
// Therefore you should always handle the error at first, then proceed to bool checks.
func Contains(src, function string) (*ContainsOutput, error) {

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

// Check represents a checker result.
type Check struct {
	Contained bool
	Matches   []*Code
	Misses    []*Code
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

// parseContains return a generic check result.
func parseContains(function string, files []string) (*Check, error) {
	out := NewCheck()

	for _, f := range files {
		c, err := Contains(f, function)

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
