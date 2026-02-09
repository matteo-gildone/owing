package owing

import (
	"flag"
	"fmt"
	"os"

	"github.com/matteo-gildone/owing/internal/parser"
)

func Main() {
	format := flag.String("format", "text", "output format: text,json,html")
	commentType := flag.String("type", "all", "comment type: TODO, FIXME, HACK, NOTE")
	exclude := flag.String("exclude", ".git,vendor,node_modules", "folders to exclude")
	flag.Parse()

	if len(flag.Args()) < 1 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage owing:\n"); err != nil {
			panic(err)
		}
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := parser.CommentParser(`// Simple patter:
// TODO: message
// FIXME: message`)

	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't parse file: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Scanning %s with format %s, type %s, exclude %s\n", flag.Arg(0), *format, *commentType, *exclude)
	os.Exit(0)
}
