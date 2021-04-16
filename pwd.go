package clubrule

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
)

// PWD returns an absolute path of caller origin.
func PWD() (string, error) {

	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return "", errors.New("failed to retrieve runtime caller")
	}

	dir, err := filepath.Abs(filepath.Dir(filename))

	if err != nil {
		return "", fmt.Errorf("failed to resolve filepath, %s", err)
	}

	return dir, nil
}
