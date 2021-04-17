package clubrule

// Contains returns true when Go source file contains the specified Go function.
// `function` parameter is case sensitive, since it checks for both exported and private functions.
// `error` will be returned in illogical situation, such as passing Non-Go files (e.g. Markdown) or file is missing.
// Therefore you should always handle the error at first, then proceed to bool checks.
func Contains(src, function string) (bool, error) {

	gots, err := Functions(src)

	if err != nil {
		return false, err
	}

	for _, f := range gots {
		if f == function {
			return true, nil
		}
	}

	return false, nil
}

// DirSourceContains returns true when source files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will not check for `*_test.go` files.
func DirSourceContains(function string, recursive bool) (bool, error) {
	files, err := Sources(recursive)

	if err != nil {
		return false, err
	}

	for _, f := range files {
		ok, err := Contains(f, function)

		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}

// DirTestContains returns true when test files in the caller directory contains the specified Go function.
// The starting directory is where you call this function.
// This function will check for `*_test.go` files only.
func DirTestContains(function string, recursive bool) (bool, error) {
	files, err := Tests(recursive)

	if err != nil {
		return false, err
	}

	for _, f := range files {
		ok, err := Contains(f, function)

		if err != nil {
			return false, err
		}

		if ok {
			return true, nil
		}
	}

	return false, nil
}
