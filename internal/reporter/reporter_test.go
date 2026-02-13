package reporter_test

import (
	"testing"

	"github.com/matteo-gildone/owing/internal/reporter"
	"github.com/matteo-gildone/owing/internal/todo"
)

func TestNewReport(t *testing.T) {
	todos := []todo.Todo{
		{File: "a.go", Type: "TODO", Line: 1, Message: "msg1"},
		{File: "a.go", Type: "FIXME", Line: 2, Message: "msg1"},
		{File: "b.go", Type: "TODO", Line: 1, Message: "msg1"},
	}

	r := reporter.NewReport(todos)

	if len(r.GroupedByFile["a.go"]) != 2 {
		t.Errorf("expected 2 todos in a.go, got %d", len(r.GroupedByFile["a.go"]))
	}

	if r.CountByType["TODO"] != 2 {
		t.Errorf("expected 2 TODOs, got %d", r.CountByType["TODO"])
	}

	if r.Total != 3 {
		t.Errorf("expected 3 total, got %d", r.Total)
	}
}
