package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/hakankaan/todo-api/pkg/logging"
)

// service is a struct for database access service
type service struct {
	l logging.Service
	r *redis.Client
}

func New(l logging.Service) (s *service) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	s = &service{
		l: l,
		r: rdb,
	}

	return
}
