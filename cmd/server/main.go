package main

import (
	"github.com/hakankaan/todo-api/pkg/http/rest"
	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/todos"
)

func main() {
	l := logging.NewStdoutLogging("DEBUG")

	rs := rest.New(l)

	ts := todos.NewService(l, rs)
	ts.InitRoutes()

	rs.RunWeb()

}
