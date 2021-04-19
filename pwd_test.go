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
