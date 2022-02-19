package todos

import (
	"github.com/hakankaan/todo-api/pkg/http/rest"
	"github.com/hakankaan/todo-api/pkg/logging"
)

type Repository interface {
	Add(string) error
	MakeDone(int64) error
	Delete(int64) error
}

type service struct {
	r  Repository
	l  logging.Service
	rs rest.RestService
}

// Newservice returns a new service for todos
func NewService(l logging.Service, rs *rest.RestService) *service {
	return &service{
		l:  l,
		rs: *rs,
	}

}

// InıtRoutes initialize routes for todos
func (ts *service) InitRoutes() {
	r := ts.rs.App

	r.Get("/ping", ts.ping)
}
