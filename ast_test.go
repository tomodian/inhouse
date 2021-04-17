package clubrule

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContainsFunctions(t *testing.T) {
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
			got, err := Functions(testfile(p.path))

			require.NoError(t, err)
			assert.Equalf(t, p.expected, len(got), spew.Sdump(p))
		}
	}

	{
		// Fail cases
		for _, p := range nonGoFiles {
			got, err := Functions(testfile(p))

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}
