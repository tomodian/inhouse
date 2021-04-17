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
func Functions(filepath string) ([]Code, error) {
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
	fileset := token.NewFileSet()

	f, err := parser.ParseFile(fileset, filepath, src, 0)

	if err != nil {
		return nil, err
	}

	outs := []Code{}

	ast.Inspect(f, func(n ast.Node) bool {
		switch fn := n.(type) {

		// Parse all function declarations.
		case *ast.FuncDecl:
			outs = append(outs, Code{
				Filepath: filepath,
				Function: fmt.Sprintf("%v", fn.Name),
				Line:     fileset.Position(fn.Pos()).Line,
			})
		}

		return true
	})

	return outs, nil
}
