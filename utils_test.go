package inhouse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	nonGoFiles = []string{
		"misc/non-existent.txt",
		"misc/binary.png",
		"misc/just-a.txt",
		"misc/README.md",
	}
)

// testfile constructs a path to test data.
// The parameter must be a relative path under ./testdata directory.
//
// e.g.
// testdata/{rel}
func testfile(rel string) string {
	return fmt.Sprintf("testdata/%s", rel)
}

func TestTestfile(t *testing.T) {
	assert.Equal(t, "testdata/hello.txt", testfile("hello.txt"))
}
