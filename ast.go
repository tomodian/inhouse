package inhouse

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

// Functions returns both exported and unexported function names.
func Functions(filepath string) ([]string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	src := string(buf)

	f, err := parser.ParseFile(token.NewFileSet(), filepath, src, 0)

	if err != nil {
		return nil, err
	}

	outs := []string{}

	ast.Inspect(f, func(n ast.Node) bool {
		switch fn := n.(type) {

		// Parse all function declarations.
		case *ast.FuncDecl:
			outs = append(outs, fmt.Sprintf("%v", fn.Name))
		}

		return true
	})

	return outs, nil
}
