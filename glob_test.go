package clubrule

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSources(t *testing.T) {
	{
		// Success cases, recursive
		type pattern struct {
			recursive bool
		}

		pats := []pattern{
			{
				recursive: true,
			},
			{
				recursive: false,
			},
		}

		for _, p := range pats {
			got, err := Sources(p.recursive)

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
		// TODO
	}
}

func TestTests(t *testing.T) {
	{
		// Success cases, recursive
		type pattern struct {
			recursive bool
		}

		pats := []pattern{
			{
				recursive: true,
			},
			{
				recursive: false,
			},
		}

		for _, p := range pats {
			got, err := Tests(p.recursive)

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
		// TODO
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

			require.NoError(t, err)
			assert.True(t, len(got) > 0)

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

			require.Error(t, err)
			assert.Nil(t, got)
		}
	}
}
