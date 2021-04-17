package inhouse

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPWD(t *testing.T) {

	got, err := PWD()

	require.NoError(t, err)
	assert.NotEmpty(t, got)
	assert.False(t, strings.HasSuffix(got, ".go"))
}
