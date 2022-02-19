package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/hakankaan/todo-api/pkg/logging"
)

type RestService struct {
	App    *fiber.App
	Logger logging.Service
}

// New creates and returns new RestService
func New(l logging.Service) *RestService {
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

func (s *RestService) RunWeb() {
	s.UseMiddlewares(logger.New(), cors.New())
	s.Logger.Error("listen", s.App.Listen(":8080"))
}
