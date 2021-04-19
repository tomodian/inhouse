package main

import (
	"testing"

	"inhouse"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("init", false)

	require.NoError(t, err)
	require.NotNil(t, got)
	assert.True(t, got.Contained)
}

func TestInternalFuctionExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("internalFunction", false)

	require.NoError(t, err)
	require.NotNil(t, got)
	assert.True(t, got.Contained)
}

func TestInternalFunctionNotExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("prohibidedFunction", false)

	require.NoError(t, err)
	assert.False(t, got.Contained)
}

func TestEnsureExportedFunctionExists(t *testing.T) {
	got, err := inhouse.SourcesContainsPWD("ExportedFunction", false)

	require.NoError(t, err)
	require.NotNil(t, got)
	assert.True(t, got.Contained)
}
