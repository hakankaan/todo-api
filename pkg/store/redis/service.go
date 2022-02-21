package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/hakankaan/todo-api/pkg/config"
	"github.com/hakankaan/todo-api/pkg/logging"
)

// service is a struct for database access service
type service struct {
	l logging.Service
	r *redis.Client
}

// New returns a new service
func New(l logging.Service, c config.Service) (s *service) {
	h := c.CFG.Redis.Host
	p := c.CFG.Redis.Port
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", h, p),
	})
	s = &service{
		l: l,
		r: rdb,
	}

	return
}
