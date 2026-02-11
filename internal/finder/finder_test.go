package finder_test

import (
	"testing"
	"testing/fstest"

	"github.com/matteo-gildone/owing/internal/finder"
)

func TestTodos(t *testing.T) {
	testFS := fstest.MapFS{
		"main.go": &fstest.MapFile{
			Data: []byte(`package main
// TODO: refactor this
func main() {}
`),
		},
		"parser.go": &fstest.MapFile{
			Data: []byte(`package parser
// FIXME: handle error
`),
		},
	}

	todos, err := finder.Todos(testFS, ".")
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) != 2 {
		t.Errorf("expected 2 todos, got %d", len(todos))
	}
}
