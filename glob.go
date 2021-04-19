package inhouse

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bmatcuk/doublestar/v3"
)

// Sources returns .go files excluding *_test.go.
func Sources(dir string, recursive bool) ([]string, error) {

	target, err := globDir(dir)

	if err != nil {
		return nil, err
	}

	return glob(globInput{
		dir:            target,
		recursive:      recursive,
		ignoreSuffixes: []string{"_test.go"},
	})
}

// Tests returns *_test.go files.
func Tests(dir string, recursive bool) ([]string, error) {

	target, err := globDir(dir)

	if err != nil {
		return nil, err
	}

	got, err := glob(globInput{
		dir:       target,
		recursive: recursive,
	})

	if err != nil {
		return nil, err
	}

	outs := []string{}

	for _, g := range got {
		if strings.HasSuffix(g, "_test.go") {
			outs = append(outs, g)
		}
	}

	return outs, nil
}

// globDir returns an absolute path of the given directory.
// Fallback to directory when given path is pointing to file.
func globDir(given string) (string, error) {
	Log(given)

	target := ""

	switch given {
	case "":
		// Use caller directory (pwd) when directory is not specified.
		pwd, err := PWD()

		if err != nil {
			return "", err
		}

		target = pwd

	default:
		home, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		for _, c := range []string{"/", "/home", home} {
			if c == given {
				return "", fmt.Errorf("directory '%s' is too broad which could bloat your machine's CPU usage", given)
			}
		}

		cwd, err := os.Getwd()

		if err != nil {
			return "", err
		}

		Log("os.Getwd", cwd)

		stat, err := os.Stat(given)

		if err != nil || os.IsNotExist(err) {
			return "", err
		}

		dir := given

		if !stat.IsDir() {
			dir = filepath.Dir(dir)
		}

		if filepath.IsAbs(dir) {
			target = dir
			break
		}

		got, err := filepath.Abs(filepath.Join(cwd, dir))

		if err != nil {
			return "", err
		}

		target = got
	}

	Log("resolved dir:", target)

	return target, nil
}

type globInput struct {
	dir            string
	recursive      bool
	ignoreSuffixes []string
}

func glob(in globInput) ([]string, error) {

	// Ensure directory exists.
	if _, err := os.Stat(in.dir); os.IsNotExist(err) {
		return nil, err
	}

	reg := ""

	if in.recursive {
		reg = fmt.Sprintf("%s/**/*.go", in.dir)
	} else {
		reg = fmt.Sprintf("%s/*.go", in.dir)
	}

	paths, err := doublestar.Glob(reg)

	if err != nil {
		return nil, err
	}

	outs := []string{}

	for _, p := range paths {
		skip := false

		for _, ig := range in.ignoreSuffixes {
			if strings.HasSuffix(p, ig) {
				skip = true
				break
			}
		}

		if !skip {
			outs = append(outs, p)
		}
	}

	return outs, nil
}
