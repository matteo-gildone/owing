package owing

import (
	"flag"
	"fmt"
	"os"
)

func Main() {
	format := flag.String("format", "json", "output format: json,html")
	commentType := flag.String("type", "all", "comment type: TODO, FIXME, HACK, NOTE")
	exclude := flag.String("exclude", ".git,vendor,node_modules", "folders to exclude")
	flag.Parse()

	if len(flag.Args()) < 1 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage owing:\n"); err != nil {
			_ = fmt.Errorf("failed to print usage: %w", err)
		}
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Scanning %s with format %s, type %s, exclude %s\n", flag.Arg(0), *format, *commentType, *exclude)
}
