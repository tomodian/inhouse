package inhouse

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

// PWD returns an absolute path of caller origin.
func PWD() (string, error) {

	current := ""

	if os.Getenv(CLIENV) == "" {
		// Comes into this context when called by source/test code.
		_, filename, _, ok := runtime.Caller(2)

		if !ok {
			return "", errors.New("failed to retrieve runtime caller")
		}

		current = filename

	} else {
		// Comes into this context when called by CLI.
		got, err := os.Getwd()

		if err != nil {
			return "", err
		}

		current = got
	}

	info, err := os.Stat(current)

	if err != nil {
		return "", err
	}

	out := ""

	if info.IsDir() {
		dir, err := filepath.Abs(current)

		if err != nil {
			return "", err
		}

		out = dir

	} else {
		dir, err := filepath.Abs(filepath.Dir(current))

		if err != nil {
			return "", err
		}

		out = dir
	}

	return out, nil
}
