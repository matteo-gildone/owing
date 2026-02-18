package todo

const (
	TypeTODO  = "TODO"
	TypeFIXME = "FIXME"
	TypeHACK  = "HACK"
	TypeNOTE  = "NOTE"
)

type Todo struct {
	File    string
	Line    int
	Type    string
	Message string
}

func NewTodo(file, todoType, message string, line int) Todo {
	return Todo{
		File:    file,
		Line:    line,
		Type:    todoType,
		Message: message,
	}
}

func FilterByType(todos []Todo, commentType string) []Todo {
	filtered := make([]Todo, 0, len(todos))

	for _, ct := range todos {
		if ct.Type == commentType {
			filtered = append(filtered, ct)
		}
	}

	return filtered
}
