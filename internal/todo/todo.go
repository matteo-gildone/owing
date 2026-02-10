package todo

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
