package clubrule

import (
	"fmt"
	"os"
	"strings"

	"github.com/bmatcuk/doublestar/v3"
)

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
