package inhouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortCode(t *testing.T) {

	got := sortCode([]*Code{
		{Filepath: "/c"},
		{Filepath: "/b"},
		{Filepath: "/a"},
	})

	assert.Equal(t, "/a", got[0].Filepath)
	assert.Equal(t, "/b", got[1].Filepath)
	assert.Equal(t, "/c", got[2].Filepath)
}
