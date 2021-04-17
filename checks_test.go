package inhouse

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCheck(t *testing.T) {
	got := NewCheck()

	assert.False(t, got.Contained)
	assert.NotNil(t, got.Matches)
	assert.NotNil(t, got.Misses)
}

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
			got, err := Contains(testfile(p.path), p.name)

			require.NoError(t, err)
			assert.Truef(t, got.HasCode(), spew.Sdump(p))
			assert.Equalf(t, p.expected, got.Contained, spew.Sdump(p))
		}
	}

	{
		// Success cases for empty Go files.
		type pattern struct {
			path string
			name string
		}

		pats := []pattern{
			{
				path: "ast/empty.go",
				name: "whatever",
			},
		}

		for _, p := range pats {
			got, err := Contains(testfile(p.path), p.name)

			require.NoError(t, err)
			assert.Falsef(t, got.Contained, spew.Sdump(p))
			assert.Falsef(t, got.HasCode(), spew.Sdump(p))
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

func TestSourcesContains(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			dir       string
			name      string
			recursive bool
			expected  bool
		}

		pats := []pattern{
			{
				dir:       "",
				name:      "SourcesContains",
				recursive: false,
				expected:  true,
			},
			{
				dir:       "./",
				name:      "SourcesContains",
				recursive: false,
				expected:  true,
			},
			{
				dir:       "./",
				name:      "SourcesContains",
				recursive: true,
				expected:  true,
			},
			{
				dir:       "",
				name:      "ExportOnly1",
				recursive: false,
				expected:  false,
			},
			{
				dir:       "",
				name:      "ExportOnly1",
				recursive: true,
				expected:  true,
			},
			{
				dir:       "./testdata",
				name:      "ExportOnly1",
				recursive: true,
				expected:  true,
			},
		}

		for _, p := range pats {
			got, err := SourcesContains(p.dir, p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, got.Contained, spew.Sdump(p), spew.Sdump(got))
		}
	}

	{
		// Fail cases
		type pattern struct {
			dir       string
			recursive bool
		}

		pats := []pattern{
			{
				dir:       "non-existent",
				recursive: false,
			},
		}

		for _, p := range pats {
			got, err := SourcesContains(p.dir, "foo", p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}

func TestTestsContains(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			dir       string
			name      string
			recursive bool
			expected  bool
		}

		pats := []pattern{
			{
				dir:       "",
				name:      "SourcesContains",
				recursive: false,
				expected:  false,
			},
			{
				dir:       "",
				name:      "TestTestsContains",
				recursive: false,
				expected:  true,
			},
		}

		for _, p := range pats {
			got, err := TestsContains(p.dir, p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, got.Contained, spew.Sdump(p))
		}
	}
}
