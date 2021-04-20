package inhouse

import (
	"os"
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

		info, err := os.Stat(got)

		require.NoError(t, err)
		assert.True(t, info.IsDir())
	}

	{
		// Success case, CLI mode
		os.Setenv(CLIENV, "yes")

		got, err := PWD()

		require.NoError(t, err)
		assert.NotEmpty(t, got)

		info, err := os.Stat(got)

		require.NoError(t, err)
		assert.True(t, info.IsDir())
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
