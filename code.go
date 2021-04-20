package inhouse

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// CodeFormat represents print styles.
type CodeFormat string

// String output.
func (c CodeFormat) String() string {
	return string(c)
}

// List of code formats.
const (
	CSVFormat   CodeFormat = "csv"
	ColonFormat CodeFormat = "colon"
	JSONFormat  CodeFormat = "json"
	TSVFormat   CodeFormat = "tsv"
)

// Code represents a Go code with some information for post processing.
type Code struct {
	Filepath string `json:"filePath"`
	Function string `json:"function"`
	Line     int    `json:"line"`

	// Just for testing.
	Tester interface{} `json:"tester,omitempty"`
}

// Format code to applicable styles.
// Suitable to be used together with CLI flags.
func (c Code) Format(style CodeFormat) string {
	switch style {

	case CSVFormat:
		return c.ToCSV()

	case JSONFormat:
		return c.ToJSON()

	case TSVFormat:
		return c.ToTSV()
	}

	return c.ToColon()
}

// ToColon is a decorator to print fields in colon-separated format.
func (c Code) ToColon() string {
	return fmt.Sprintf("%s:%s:%d", c.Filepath, c.Function, c.Line)
}

// ToCSV is a decorator to print fields in comma-separated format.
func (c Code) ToCSV() string {
	return fmt.Sprintf("%s,%s,%d", c.Filepath, c.Function, c.Line)
}

// ToTSV is a decorator to print fields in tab-separated format.
func (c Code) ToTSV() string {
	return fmt.Sprintf("%s\t%s\t%d", c.Filepath, c.Function, c.Line)
}

// ToJSON is a decorator to print fields in JSON format.
// Note this method will exit immediately on failure.
func (c Code) ToJSON() string {

	b, err := json.Marshal(c)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(b)
}

// Sort slice of code by filename.
func sortCode(given []*Code) []*Code {
	sort.Slice(given[:], func(i, j int) bool {
		return given[i].Filepath < given[j].Filepath
	})

	return given
}
