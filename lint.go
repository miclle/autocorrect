package autocorrect

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
)

func main() {
	filename := "_fixtures/example.go"
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println(Lint(string(raw)))
}

// Lint for format Go source file to auto correct String / Comment
func Lint(raw string) (out string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "", string(raw), parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		var s string
		switch x := n.(type) {
		case *ast.BasicLit:
			s = x.Value
			x.Value = Format(s)
		case *ast.Comment:
			x.Text = Format(x.Text)
		}

		return true
	})

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, node); err != nil {
		panic(err)
	}

	out = string(buf.Bytes())
	return out
}
