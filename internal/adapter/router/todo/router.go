package todo

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"mongmeo.dev/todo/internal/application/todo"
	"mongmeo.dev/todo/internal/pkg/er"
	"net/http"
)

type Router interface {
	GetRouter() chi.Router
}

type router struct {
	r               chi.Router
	todoApplication todo.Application
}

func (router *router) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todoList, err := router.todoApplication.GetTodoList(r.Context())
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		handleError(w, err)
		return
	}
	marshal, err := json.Marshal(todoList)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(200)
	w.Write(marshal)
}

func (router *router) get(path string, handlerFunc http.HandlerFunc) {
	router.GetRouter().Get(path, handlerFunc)
}

func (router *router) GetRouter() chi.Router {
	return router.r
}

func New(todoApplication todo.Application) Router {
	r := chi.NewRouter()
	router := &router{r: r, todoApplication: todoApplication}
	router.get("/", router.GetAllTodos)
	return router
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(er.GetResponseStatus(err))
	response, _ := er.HelperErrorResponse(err)
	w.Write(response)
}
