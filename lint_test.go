package autocorrect

import (
	"testing"
)

func TestLint(t *testing.T) {
	raw := readFile("example.go")
	expected := readFile("example.go.expected")
	out := Lint(raw)

	if out != expected {
		t.Errorf("\n== Expected ==========\n\n%s\n\n== Actual ==========\n\n%s", expected, out)
	}
}
