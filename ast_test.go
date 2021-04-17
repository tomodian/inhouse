package clubrule

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFunctions(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			path     string
			expected int
		}

		pats := []pattern{
			{
				path:     "ast/empty.go",
				expected: 0,
			},
			{
				path:     "ast/exportonly.go",
				expected: 2,
			},
			{
				path:     "ast/mix.go",
				expected: 2,
			},
			{
				path:     "ast/privateonly.go",
				expected: 2,
			},
		}

		for _, p := range pats {
			got, err := Functions(fmt.Sprintf("testdata/%s", p.path))

			require.NoError(t, err)
			assert.Equalf(t, p.expected, len(got), spew.Sdump(p))
		}
	}

	{
		// Fail cases
		pats := []string{
			"misc/non-existent.txt",
			"misc/binary.png",
			"misc/just-a.txt",
			"misc/README.md",
		}

		for _, p := range pats {
			got, err := Functions(fmt.Sprintf("testdata/%s", p))

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}
