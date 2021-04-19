package inhouse

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// PWD returns an absolute path of caller origin.
func PWD() (string, error) {

	cwd := ""

	if os.Getenv(CLIENV) == "" {
		// Comes into this context when called by source/test code.
		_, filename, _, ok := runtime.Caller(2)

		if !ok {
			return "", errors.New("failed to retrieve runtime caller")
		}

		cwd = filename

	} else {
		// Comes into this context when called by CLI.
		got, err := os.Getwd()

		if err != nil {
			return "", err
		}

		cwd = got
	}

	dir, err := filepath.Abs(filepath.Dir(cwd))

	if err != nil {
		return "", fmt.Errorf("failed to resolve filepath, %s", err)
	}

	return dir, nil
}
