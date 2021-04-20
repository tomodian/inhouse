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

	return parseDir(current)
}

func parseDir(given string) (string, error) {
	info, err := os.Stat(given)

	if err != nil {
		return "", err
	}

	if info.IsDir() {
		dir, err := filepath.Abs(given)

		if err != nil {
			return "", err
		}

		return dir, nil
	}

	dir, err := filepath.Abs(filepath.Dir(given))

	if err != nil {
		return "", err
	}

	return dir, nil
}
