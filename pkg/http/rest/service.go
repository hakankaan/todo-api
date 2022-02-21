package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hakankaan/todo-api/pkg/config"
	"github.com/hakankaan/todo-api/pkg/logging"
)

const defaultPort = "8080"

type RestService struct {
	App    *fiber.App
	Logger logging.Service
	Config config.Config
}

// New creates and returns new RestService
func New(l logging.Service, c config.Service) *RestService {
	app := fiber.New()
	return &RestService{
		App:    app,
		Logger: l,
	}
}

// UseMiddlewares takes middlewares as argument and adds middlewares to fiber.App
func (s *RestService) UseMiddlewares(middlewares ...func(*fiber.Ctx) error) {
	for _, mw := range middlewares {
		s.App.Use(mw)
	}
}

// RunWeb starts fiber.App
func (s *RestService) RunWeb() {
	p := s.Config.Http.Port
	if p == "" {
		p = defaultPort
	}

	s.UseMiddlewares(logger.New(), cors.New())
	s.Logger.Error("listen", s.App.Listen(fmt.Sprintf(":%s", p)))
}
