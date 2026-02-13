package formatter

import (
	"fmt"
	"io"
	"sort"

	"github.com/matteo-gildone/owing/internal/reporter"
	"github.com/matteo-gildone/owing/internal/styles"
	"github.com/matteo-gildone/owing/internal/todo"
)

func Text(w io.Writer, r reporter.Report) error {
	baseStyle := styles.NewStyles()
	fileStyle := baseStyle.Bold()
	dimStyle := baseStyle.Dim()
	fmt.Fprintf(w, "Found %d TODOs in %d files\n", r.Total, len(r.GroupedByFile))

	types := make([]string, 0, len(r.CountByType))

	for typ := range r.CountByType {
		types = append(types, typ)
	}

	sort.Strings(types)
	for _, typ := range types {
		typeStyle := getStyleForType(typ)
		fmt.Fprintf(w, "%s: %d   ", typeStyle.Render(fmt.Sprintf("%s", typ)), r.CountByType[typ])
	}

	fmt.Fprintln(w)

	files := make([]string, 0, len(r.GroupedByFile))

	for file := range r.GroupedByFile {
		files = append(files, file)
	}

	sort.Strings(files)

	for _, file := range files {
		todos := r.GroupedByFile[file]
		header := fmt.Sprintf("%s (%d):\n", file, len(todos))
		fmt.Fprint(w, fileStyle.Render(header))
		for _, t := range todos {
			typeStyle := getStyleForType(t.Type)
			fmt.Fprintf(w, "  %s %s %s\n", dimStyle.Render(fmt.Sprintf("%-4d", t.Line)), typeStyle.Render(fmt.Sprintf("[%s]", t.Type)), t.Message)
		}
		fmt.Fprintln(w)
	}

	return nil
}

func getStyleForType(todoType string) styles.Style {
	base := styles.NewStyles()

	switch todoType {
	case todo.TypeTODO:
		return base.Cyan()
	case todo.TypeFIXME:
		return base.Yellow()
	case todo.TypeHACK:
		return base.Red()
	case todo.TypeNOTE:
		return base.Green()
	default:
		return base
	}
}
