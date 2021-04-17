package inhouse

import (
	"bytes"
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

		Log("hello")
		w.Close()

		var buf bytes.Buffer
		buf.ReadFrom(r)

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

		expected := "hello"
		Log(expected)
		w.Close()

		var buf bytes.Buffer
		buf.ReadFrom(r)

		assert.Equal(t, expected, strings.TrimRight(buf.String(), "\n"))

		// Teardown
		os.Setenv(debug, orig)
	}
}
