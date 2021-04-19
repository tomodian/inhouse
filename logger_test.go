package inhouse

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	stdout := os.Stdout

	defer func() {
		os.Stdout = stdout
	}()

	{
		// Ensure log message does not appear when DEBUG environment variable is not set.
		r, w, err := os.Pipe()
		require.NoError(t, err)

		os.Stdout = w

		orig := os.Getenv(debug)
		os.Unsetenv(debug)

		Log("hello", "world")
		w.Close()

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		require.NoError(t, err)

		assert.Equal(t, "", strings.TrimRight(buf.String(), "\n"))

		// Teardown
		os.Setenv(debug, orig)
	}

	{
		// Ensure log message appears with DEBUG=1
		r, w, err := os.Pipe()
		require.NoError(t, err)

		os.Stdout = w

		orig := os.Getenv(debug)
		os.Setenv(debug, "1")

		sample := "hello"
		Log(sample)
		w.Close()

		var buf bytes.Buffer
		_, err = buf.ReadFrom(r)
		require.NoError(t, err)

		exp := fmt.Sprintf("%s: %s", debug, sample)
		got := strings.TrimRight(buf.String(), "\n")

		assert.Equal(t, exp, got)

		// Teardown
		os.Setenv(debug, orig)
	}
}
