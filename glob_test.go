package inhouse

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSources(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			dir       string
			recursive bool
		}

		pats := []pattern{
			{
				dir:       "",
				recursive: true,
			},
			{
				dir:       "",
				recursive: false,
			},
			{
				dir:       "./",
				recursive: false,
			},
			{
				dir:       "./testdata",
				recursive: false,
			},
			{
				dir:       "./.github",
				recursive: false,
			},
		}

		for _, p := range pats {
			got, err := Sources(p.dir, p.recursive)

			require.NoError(t, err)
			assert.NotNil(t, got)

			for _, g := range got {
				name := filepath.Base(g)

				assert.Falsef(t, strings.HasSuffix(name, "_test.go"), spew.Sdump(p, got))
			}
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
			{
				dir:       "non-existent",
				recursive: true,
			},
		}

		for _, p := range pats {
			got, err := Sources(p.dir, p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}

func TestTests(t *testing.T) {
	{
		// Success cases, recursive
		type pattern struct {
			dir       string
			recursive bool
		}

		pats := []pattern{
			{
				dir:       "",
				recursive: true,
			},
			{
				dir:       "",
				recursive: false,
			},
			{
				dir:       "./",
				recursive: true,
			},
		}

		for _, p := range pats {
			got, err := Tests(p.dir, p.recursive)

			require.NoError(t, err)
			assert.NotNil(t, got)

			for _, g := range got {
				name := filepath.Base(g)

				assert.Truef(t, strings.HasSuffix(name, "_test.go"), spew.Sdump(p, got))
			}
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
			{
				dir:       "non-existent",
				recursive: true,
			},
		}

		for _, p := range pats {
			got, err := Tests(p.dir, p.recursive)

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}
func TestGlobDir(t *testing.T) {
	{
		// Success cases
		type pattern struct {
			dir string
		}

		pats := []pattern{
			{dir: ""},
			{dir: "testdata"},
			{dir: "./testdata"},
			// Fallback to dir.
			{dir: "./testdata/misc/just-a.txt"},
		}

		cwd, err := os.Getwd()
		require.NoError(t, err)

		for _, p := range pats {
			got, err := globDir(p.dir)

			require.NoErrorf(t, err, spew.Sdump(p))

			assert.Truef(t, strings.HasPrefix(got, cwd), spew.Sdump(got))
			assert.NotSamef(t, p.dir, got, spew.Sdump(p))
		}
	}

	{
		// Fail cases
		type pattern struct {
			dir string
		}

		pats := []pattern{
			{dir: "non-existent"},
		}

		for _, p := range pats {
			got, err := globDir(p.dir)

			require.Errorf(t, err, spew.Sdump(p))
			assert.Emptyf(t, got, spew.Sdump(p))
		}
	}
}

func TestGlob(t *testing.T) {

	dir, err := PWD()
	require.NoError(t, err)

	{
		// Success cases
		pats := []globInput{
			{
				dir: dir,
			},
			{
				dir: dir,
			},
			{
				dir:            dir,
				ignoreSuffixes: []string{"_test.go"},
			},
			{
				dir:            dir,
				ignoreSuffixes: []string{"ensure.go"},
			},
		}

		for _, p := range pats {
			got, err := glob(p)

			require.NoErrorf(t, err, spew.Sdump(p))
			assert.Truef(t, len(got) > 0, spew.Sdump(p))

			for _, g := range got {
				name := filepath.Base(g)

				for _, s := range p.ignoreSuffixes {
					assert.Falsef(t, strings.HasSuffix(name, s), spew.Sdump(p, got))
				}
			}
		}
	}

	{
		// Fail cases
		pats := []globInput{
			{
				dir: "non-existent-directory",
			},
		}

		for _, p := range pats {
			got, err := glob(p)

			require.Errorf(t, err, spew.Sdump(p))
			assert.Nilf(t, got, spew.Sdump(p))
		}
	}
}
