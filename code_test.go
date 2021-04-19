package inhouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodeFormatString(t *testing.T) {
	assert.Equal(t, string(CSVFormat), CSVFormat.String())
}

func TestCodeDecorators(t *testing.T) {
	c := Code{
		Filepath: "/a",
		Function: "b",
		Line:     100,
	}

	var (
		colon = "/a:b:100"
		csv   = "/a,b,100"
		tsv   = "/a\tb\t100"
		json  = `{"filePath":"/a","function":"b","line":100}`
	)

	// Check raw outputs.
	assert.Equal(t, colon, c.ToColon())
	assert.Equal(t, csv, c.ToCSV())
	assert.Equal(t, tsv, c.ToTSV())
	assert.Equal(t, json, c.ToJSON())

	// Check wrapper results are exactly the same.
	assert.Equal(t, c.ToColon(), c.Format(ColonFormat))
	assert.Equal(t, c.ToCSV(), c.Format(CSVFormat))
	assert.Equal(t, c.ToTSV(), c.Format(TSVFormat))
	assert.Equal(t, c.ToJSON(), c.Format(JSONFormat))
}

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
