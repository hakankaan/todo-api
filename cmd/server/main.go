package main

import (
	"github.com/hakankaan/todo-api/pkg/config"
	"github.com/hakankaan/todo-api/pkg/http/rest"
	"github.com/hakankaan/todo-api/pkg/logging"
	"github.com/hakankaan/todo-api/pkg/store/postgres"
	"github.com/hakankaan/todo-api/pkg/store/redis"
	"github.com/hakankaan/todo-api/pkg/todos"
)

func main() {

	c := config.New()

	l := logging.NewStdoutLogging(c.CFG.Logging.Level)

	rs := rest.New(l, c)

	ps := postgres.New(l, c)

	redis := redis.New(l, c)

	ts := todos.NewService(l, rs, ps, redis)
	ts.InitRoutes()

	rs.RunWeb()

}
