package main

import (
	"os"

	"github.com/hakankaan/todo-api/pkg/http/rest"
	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/store/postgres"
	"github.com/hakankaan/todo-api/pkg/todos"
)

func main() {
	l := logging.NewStdoutLogging("DEBUG")

	rs := rest.New(l)

	ps, err := postgres.New(l)
	if err != nil {
		l.Error("postgres.New", err)
		os.Exit(1)
	}

	ts := todos.NewService(l, rs, ps)
	ts.InitRoutes()

	rs.RunWeb()

}
