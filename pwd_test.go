package inhouse

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPWD(t *testing.T) {
	{
		// Success case, source mode
		got, err := PWD()

		require.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.False(t, strings.HasSuffix(got, ".go"))
	}

	{
		// Success case, CLI mode
		os.Setenv(CLIENV, "yes")

		got, err := PWD()

		require.NoError(t, err)
		assert.NotEmpty(t, got)
		assert.False(t, strings.HasSuffix(got, ".go"))
	}
}

func TestParseDir(t *testing.T) {
	{
		// Success case
		pats := []string{
			"./",
			"./testdata",
		}

		for _, p := range pats {
			got, err := parseDir(p)

			require.NoError(t, err)
			assert.NotEmpty(t, got)
		}
	}

	{
		// Fail cases
		pats := []string{
			"./non-existent",
		}

		for _, p := range pats {
			got, err := parseDir(p)

			require.Error(t, err)
			assert.Empty(t, got)
		}
	}
}
