package todo

import "testing"

func TestFilterByType(t *testing.T) {
	todos := []Todo{
		{File: "a.go", Type: "TODO", Line: 1, Message: "msg1"},
		{File: "a.go", Type: "FIXME", Line: 2, Message: "msg1"},
		{File: "b.go", Type: "TODO", Line: 1, Message: "msg1"},
	}

	filtered := FilterByType(todos, "TODO")

	if len(filtered) != 2 {
		t.Errorf("expected 2 TODOs, got: %d", len(filtered))
	}
}
