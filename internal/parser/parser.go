package parser

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`(TODO|FIXME|HACK|NOTE):\s*(.+)`)

type Match struct {
	Type    string
	Message string
	Line    int
}

func FileParser(text string) []Match {
	var todos []Match
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanLines)
	var line int
	for scanner.Scan() {
		matches := pattern.FindStringSubmatch(scanner.Text())

		if len(matches) > 2 {
			todos = append(todos, Match{Type: matches[1], Message: matches[2], Line: line + 1})
			fmt.Printf("type: %q, message: %q, line:%d\n", matches[1], matches[2], line+1)
		}

		line++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	fmt.Printf("%#v\n", todos)

	return todos
}
