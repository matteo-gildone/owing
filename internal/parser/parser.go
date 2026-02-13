package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/matteo-gildone/owing/internal/todo"
)

var regex = fmt.Sprintf("(%s|%s|%s|%s):\\s*(.+)", todo.TypeTODO, todo.TypeFIXME, todo.TypeHACK, todo.TypeNOTE)
var pattern = regexp.MustCompile(regex)

type Match struct {
	Type    string
	Message string
	Line    int
}

func CommentParser(text string) ([]Match, error) {
	var todos []Match
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanLines)
	var line int
	for scanner.Scan() {
		line++
		if matches := pattern.FindStringSubmatch(scanner.Text()); len(matches) > 2 {
			todos = append(todos, Match{Type: matches[1], Message: matches[2], Line: line})
		}

	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading standard input: %w", err)
	}

	return todos, nil
}
