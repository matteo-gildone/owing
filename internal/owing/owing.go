package owing

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/matteo-gildone/owing/internal/finder"
	"github.com/matteo-gildone/owing/internal/formatter"
	"github.com/matteo-gildone/owing/internal/reporter"
	"github.com/matteo-gildone/owing/internal/todo"
)

func Main() {
	format := flag.String("format", "text", "output format: text,json,html")
	commentType := flag.String("type", "all", "comment type: TODO, FIXME, HACK, NOTE")
	exclude := flag.String("exclude", ".git,vendor,node_modules", "folders to exclude")
	flag.Parse()

	dir := flag.Arg(0)

	excludeFoldersDefaults := []string{".git", "vendor", "node_modules", "testdata", "script"}
	excludeFolders := make(map[string]struct{})

	for _, d := range excludeFoldersDefaults {
		excludeFolders[d] = struct{}{}
	}

	for _, folder := range strings.Split(*exclude, ",") {
		trimmed := strings.TrimSpace(folder)
		if trimmed != "" {
			excludeFolders[trimmed] = struct{}{}
		}
	}

	if len(flag.Args()) < 1 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage owing:\n"); err != nil {
			panic(err)
		}
		flag.PrintDefaults()
		os.Exit(1)
	}

	fsys := os.DirFS(dir)

	todos, err := finder.Todos(fsys, ".", excludeFolders)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't parse files: %v", err)
		os.Exit(1)
	}

	if *commentType != "all" {
		todos = todo.FilterByType(todos, *commentType)
	}

	report := reporter.NewReport(todos)

	switch *format {
	case "html":
		err = formatter.Html(os.Stdout, report)
	default:
		err = formatter.Text(os.Stdout, report)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't format files: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}
