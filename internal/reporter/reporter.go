package reporter

import "github.com/matteo-gildone/owing/internal/todo"

type Report struct {
	CountByType   map[string]int
	GroupedByFile map[string][]todo.Todo
	Total         int
}

func NewReport(todos []todo.Todo) Report {
	r := Report{
		CountByType:   make(map[string]int),
		GroupedByFile: make(map[string][]todo.Todo),
	}

	for _, t := range todos {
		r.GroupedByFile[t.File] = append(r.GroupedByFile[t.File], t)
		r.CountByType[t.Type]++
		r.Total++
	}

	return r
}
