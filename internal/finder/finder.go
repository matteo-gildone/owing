package finder

import (
	"fmt"
	"io/fs"

	"github.com/matteo-gildone/owing/internal/parser"
	"github.com/matteo-gildone/owing/internal/todo"
)

func Todos(fsys fs.FS, root string, exclude map[string]struct{}) ([]todo.Todo, error) {
	var todos []todo.Todo
	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			if _, ok := exclude[d.Name()]; ok {
				return fs.SkipDir
			}
			return nil
		}

		if !d.Type().IsRegular() {
			return nil
		}

		content, err := fs.ReadFile(fsys, path)

		if err != nil {
			return fmt.Errorf("failed reading file: %w", err)
		}

		matches, err := parser.CommentParser(string(content))

		if err != nil {
			return fmt.Errorf("failed to parse content: %w", err)
		}

		for _, match := range matches {
			todos = append(todos, todo.NewTodo(path, match.Type, match.Message, match.Line))
		}

		return nil
	})
	return todos, err
}
