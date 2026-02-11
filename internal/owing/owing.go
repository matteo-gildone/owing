package owing

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/owing/internal/finder"
)

func Main() {
	//format := flag.String("format", "text", "output format: text,json,html")
	//commentType := flag.String("type", "all", "comment type: TODO, FIXME, HACK, NOTE")
	//exclude := flag.String("exclude", ".git,vendor,node_modules", "folders to exclude")
	flag.Parse()

	dir := flag.Arg(0)

	if len(flag.Args()) < 1 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage owing:\n"); err != nil {
			panic(err)
		}
		flag.PrintDefaults()
		os.Exit(1)
	}

	fsys := os.DirFS(dir)

	todos, err := finder.Files(fsys)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't parse files: %v", err)
		os.Exit(1)
	}

	for _, todo := range todos {
		fmt.Printf("%s:%d [%s] %s\n", todo.File, todo.Line, todo.Type, todo.Message)
	}
	os.Exit(0)
}
