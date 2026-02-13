package formatter

import (
	"fmt"
	"io"
	"sort"

	"github.com/matteo-gildone/owing/internal/reporter"
)

func Text(w io.Writer, r reporter.Report) error {
	fmt.Fprintf(w, "Found %d TODOs in %d files\n", r.Total, len(r.GroupedByFile))

	types := make([]string, 0, len(r.CountByType))

	for typ := range r.CountByType {
		types = append(types, typ)
	}

	sort.Strings(types)

	for _, typ := range types {
		fmt.Fprintf(w, "  %s: %d\n", typ, r.CountByType[typ])
	}

	fmt.Fprintln(w)

	files := make([]string, 0, len(r.GroupedByFile))

	for file := range r.GroupedByFile {
		files = append(files, file)
	}

	sort.Strings(files)

	for _, file := range files {
		todos := r.GroupedByFile[file]
		fmt.Fprintf(w, "%s (%d):\n", file, len(todos))
		for _, t := range todos {
			fmt.Fprintf(w, "  %d [%s] %s\n", t.Line, t.Type, t.Message)
		}
		fmt.Fprintln(w)
	}

	return nil
}
