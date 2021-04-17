package clubrule

// ContainsFunction returns true when Go source file contains the specified function.
// `function` parameter is case sensitive, since it checks for both exported and private functions.
// `error` will be returned in illogical situation, such as passing Non-Go files (e.g. Markdown) or file is missing.
// Therefore you should always handle the error at first, then proceed to bool checks.
func ContainsFunction(src, function string) (bool, error) {

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
