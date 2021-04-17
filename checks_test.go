package clubrule

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContainsFunction(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			path     string
			name     string
			expected bool
		}

		pats := []pattern{
			{
				path:     "ast/empty.go",
				name:     "whatever",
				expected: false,
			},
			{
				path:     "ast/init.go",
				name:     "init",
				expected: true,
			},
			{
				path:     "ast/exportonly.go",
				name:     "ExportOnly1",
				expected: true,
			},
			{
				path:     "ast/privateonly.go",
				name:     "privateOnly1",
				expected: true,
			},
		}

		for _, p := range pats {
			ok, err := ContainsFunction(testfile(p.path), p.name)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, ok, spew.Sdump(p))
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
