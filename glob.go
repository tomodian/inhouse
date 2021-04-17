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

func globDir(dir string) (string, error) {
	target := ""

	switch dir {
	case "":
		// Use caller directory (pwd) when directory is not specified.
		pwd, err := PWD()

		if err != nil {
			return "", err
		}

		target = pwd

	default:
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return "", err
		}

		got, err := filepath.Abs(filepath.Dir(dir))

		if err != nil {
			return "", err
		}

		target = got
	}

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
