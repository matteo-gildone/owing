package formatter_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/matteo-gildone/owing/internal/formatter"
	"github.com/matteo-gildone/owing/internal/reporter"
	"github.com/matteo-gildone/owing/internal/todo"
)

func TestText(t *testing.T) {
	report := reporter.Report{
		GroupedByFile: map[string][]todo.Todo{
			"a.go": {
				{Type: "TODO", Line: 1, Message: "fix this"},
				{Type: "FIXME", Line: 2, Message: "refactor"},
			},
			"b.go": {
				{Type: "HACK", Line: 1, Message: "temporary"},
			},
		},
		CountByType: map[string]int{
			"TODO":  1,
			"FIXME": 1,
			"HACK":  1,
		},
		Total: 3,
	}

	var buf bytes.Buffer
	err := formatter.Text(&buf, report)
	if err != nil {
		t.Fatal(err)
	}

	output := buf.String()

	if !strings.Contains(output, "Found 3 TODOs in 2 files") {
		t.Errorf("missing header in output:\n%s", output)
	}

	if !strings.Contains(output, "TODO: 1") {
		t.Errorf("missing TODO stats in output:\n%s", output)
	}

	if !strings.Contains(output, "a.go (2):") {
		t.Errorf("missing file group in output:\n%s", output)
	}
}
