package todo

import (
	"context"
	"log"
	"mongmeo.dev/todo/internal/adapter/database/ent"
	TodoModel "mongmeo.dev/todo/internal/application/todo/model"
	"mongmeo.dev/todo/internal/pkg/er"
)

type Application interface {
	GetTodoList(ctx context.Context) ([]TodoModel.TodoResponse, error)
}

type application struct {
	todoClient *ent.TodoClient
}

func (a *application) GetTodoList(ctx context.Context) ([]TodoModel.TodoResponse, error) {
	list, err := a.todoClient.Query().All(ctx)
	if err != nil {
		log.Println(err)
		return nil, er.WithCode(err, er.CodeUnspecified)
	}

	response := make([]TodoModel.TodoResponse, 0)

	for _, todo := range list {
		response = append(response, TodoModel.TodoResponse{
			Title:  todo.Title,
			IsDone: todo.IsDone,
		})
	}

	return response, nil
}

func New(todoClient *ent.TodoClient) Application {
	return &application{todoClient: todoClient}
}
