package todos

import (
	"github.com/hakankaan/todo-api/pkg/http/rest"
	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/store"
)

type Repository interface {
	GetTodo(string) (store.Todo, error)
	AddTodo(store.Todo) error
	MakeDone(string) error
	DeleteTodo(string) error
}

type service struct {
	r  Repository
	l  logging.Service
	rs rest.RestService
}

// Newservice returns a new service for todos
func NewService(l logging.Service, rs *rest.RestService, r Repository) *service {
	return &service{
		l:  l,
		rs: *rs,
		r:  r,
	}
}

// InıtRoutes initialize routes for todos
func (ts *service) InitRoutes() {
	r := ts.rs.App

	r.Get("/ping", ts.ping)
	r.Post("/api/todos", ts.addTodo)
	r.Get("/api/todos/:guid", ts.getTodo)
	r.Put("/api/todos/:guid", ts.makeDone)
	r.Delete("/api/todos/:guid", ts.deleteTodo)
}
