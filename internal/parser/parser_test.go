package parser_test

import (
	"testing"

	"github.com/matteo-gildone/owing/internal/parser"
)

func TestFileParser(t *testing.T) {
	t.Run("expect not todos", func(t *testing.T) {
		got, err := parser.CommentParser(`// Simple patter:
// no todo
`)
		if err != nil {
			t.Fatal(err)
		}

		if len(got) != 0 {
			t.Error("expected no todos")
		}
	})

	t.Run("expect find matches", func(t *testing.T) {
		got, err := parser.CommentParser(`// Simple patter:
// TODO: message
// FIXME: message`)

		if err != nil {
			t.Fatal(err)
		}

		if len(got) != 2 {
			t.Errorf("expected 2 todos, got: %d", len(got))
		}

		if got[0].Type != "TODO" {
			t.Errorf("expected type: TODO, got: %q", got[0].Type)
		}
	})

	t.Run("expect find matches with different comment style", func(t *testing.T) {
		got, err := parser.CommentParser(`# Simple patter:
# TODO: message
# FIXME: message`)

		if err != nil {
			t.Fatal(err)
		}

		if len(got) != 2 {
			t.Errorf("expected 2 todos, got: %d", len(got))
		}

		if got[0].Type != "TODO" {
			t.Errorf("expected type: TODO, got: %q", got[0].Type)
		}
	})
}
