package inhouse

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
}

func TestCodeToJSON(t *testing.T) {
	c := Code{
		Filepath: "/a",
		Function: "b",
		Line:     100,
	}

	{
		// Success case
		assert.Equal(t, c.ToJSON(), c.Format(JSONFormat))
	}

	{
		// Fail case
		if os.Getenv("BE_CRASHER") == "1" {
			// Uses this tips to make json.Marshal to fail.
			// https://stackoverflow.com/a/48901259/515244
			c.Tester = make(chan int)

			c.ToJSON()
			return
		}

		cmd := exec.Command(os.Args[0], "-test.run=TestCodeToJSON")
		cmd.Env = append(os.Environ(), "BE_CRASHER=1")

		err := cmd.Run()
		res, ok := err.(*exec.ExitError)

		require.True(t, ok)
		require.False(t, res.Success())
	}
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
