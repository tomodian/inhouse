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

func TestCheckCombine(t *testing.T) {
	d := NewCheck()

	d.Matches = []*Code{
		{Filepath: "/b", Function: "b1", Line: 1},
		{Filepath: "/b", Function: "b2", Line: 2},
	}

	d.Misses = []*Code{
		{Filepath: "/a", Function: "a1", Line: 1},
		{Filepath: "/a", Function: "a2", Line: 2},
	}

	got := d.Combine()

	assert.Equal(t, len(d.Matches)+len(d.Misses), len(got))
	assert.Equal(t, "/a", got[0].Filepath)
}

func TestCheckToJSON(t *testing.T) {
	{
		// Success case
		d := NewCheck()

		got, err := d.ToJSON()

		require.NoError(t, err)
		assert.NotEmpty(t, got)
	}

	{
		// Fail case
		d := NewCheck()

		// Uses this tips to make json.Marshal to fail.
		// https://stackoverflow.com/a/48901259/515244
		d.Tester = make(chan int)

		got, err := d.ToJSON()

		require.Error(t, err)
		assert.Empty(t, got)
	}
}

func TestMatchFunction(t *testing.T) {
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
			got, err := matchFunction(testfile(p.path), p.name)

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
			got, err := matchFunction(testfile(p.path), p.name)

			require.NoError(t, err)
			assert.Falsef(t, got.Contained, spew.Sdump(p))
			assert.Falsef(t, got.HasCode(), spew.Sdump(p))
		}
	}

	{
		// Fail cases
		for _, p := range nonGoFiles {
			_, err := matchFunction(testfile(p), "whatever")

			require.Error(t, err)
		}
	}
}

func TestContains(t *testing.T) {
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
			{
				dir:       "",
				name:      "TestTestsContains",
				recursive: true,
				expected:  true,
			},
			{
				dir:       "examples",
				name:      "TestInitExists",
				recursive: true,
				expected:  true,
			},
		}

		for _, p := range pats {
			got, err := Contains(p.dir, p.name, p.recursive)

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
			got, err := Contains(p.dir, "foo", p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}

func TestContainsPWD(t *testing.T) {
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
				name:      "SourcesContains",
				recursive: false,
				expected:  true,
			},
			{
				name:      "SourcesContains",
				recursive: true,
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
			{
				name:      "ExportOnly1",
				recursive: true,
				expected:  true,
			},
			{
				name:      "TestTestsContains",
				recursive: true,
				expected:  true,
			},
			{
				name:      "TestInitExists",
				recursive: true,
				expected:  true,
			},
		}

		for _, p := range pats {
			got, err := ContainsPWD(p.name, p.recursive)

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
			got, err := Contains(p.dir, "foo", p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
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

func TestSourcesContainsPWD(t *testing.T) {
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
				name:      "SourcesContains",
				recursive: true,
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
			{
				name:      "ExportOnly1",
				recursive: true,
				expected:  true,
			},
		}

		for _, p := range pats {
			got, err := SourcesContainsPWD(p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, got.Contained, spew.Sdump(p), spew.Sdump(got))
		}
	}

	{
		// Fail cases
		type pattern struct {
			recursive bool
		}

		pats := []pattern{
			{
				recursive: false,
			},
			{
				recursive: true,
			},
		}

		for _, p := range pats {
			got, err := SourcesContainsPWD("foo", p.recursive)

			require.NoErrorf(t, err, spew.Sdump(p))
			require.NotNilf(t, got, spew.Sdump(p))
			assert.Falsef(t, got.Contained, spew.Sdump(p))
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
			got, err := TestsContains(p.dir, "foo", p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}

func TestTestsContainsPWD(t *testing.T) {
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
			got, err := TestsContainsPWD(p.name, p.recursive)

			require.NoError(t, err)
			assert.Equalf(t, p.expected, got.Contained, spew.Sdump(p))
		}
	}

	{
		// Fail cases
		type pattern struct {
			recursive bool
		}

		pats := []pattern{
			{
				recursive: false,
			},
			{
				recursive: true,
			},
		}

		for _, p := range pats {
			got, err := TestsContainsPWD("foo", p.recursive)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.NotNilf(t, got, spew.Sdump(p))
			assert.False(t, got.Contained)
		}
	}
}
