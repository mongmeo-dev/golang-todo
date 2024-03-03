package golang_todo

import "mongmeo.dev/todo/internal/pkg/er"

var (
	ErrBadRequest = er.New("golang-todo.errors: BadGateway")
)
