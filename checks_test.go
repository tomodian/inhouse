package clubrule

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
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
			ok, err := Contains(testfile(p.path), p.name)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, ok, spew.Sdump(p))
		}
	}

	{
		// Fail cases
		for _, p := range nonGoFiles {
			_, err := Contains(testfile(p), "whatever")

			require.Error(t, err)
		}
	}
}

func TestSourcesHas(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			name      string
			recursive bool
			expected  bool
		}

		pats := []pattern{
			{
				name:      "SourcesContains",
				recursive: false,
				expected:  true,
			},
			{
				name:      "ExportOnly1",
				recursive: false,
				expected:  false,
			},
			{
				name:      "ExportOnly1",
				recursive: true,
				expected:  true,
			},
		}

		for _, p := range pats {
			ok, err := SourcesContains(p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, ok, spew.Sdump(p))
		}
	}
}

func TestTestsContains(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			name      string
			recursive bool
			expected  bool
		}

		pats := []pattern{
			{
				name:      "SourcesContains",
				recursive: false,
				expected:  false,
			},
			{
				name:      "TestTestsContains",
				recursive: false,
				expected:  true,
			},
		}

		for _, p := range pats {
			ok, err := TestsContains(p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, ok, spew.Sdump(p))
		}
	}
}
