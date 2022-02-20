package todos

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hakankaan/todo-api/pkg/store"
)

type baseResponse struct {
	Msg string `json:"message"`
}

type todoResponse struct {
	baseResponse
	store.Todo `json:"data"`
}

func (ts *service) ping(c *fiber.Ctx) error {
	return c.SendString("success")
}

// addTodo validates request with AddTodoRequest then adds todo to store
func (ts *service) addTodo(c *fiber.Ctx) error {
	var r AddTodoRequest
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	guid := uuid.New().String()

	t := store.Todo{
		GuID:        guid,
		Title:       r.Title,
		Description: r.Description,
	}
	err := ts.r.AddTodo(t)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	resp := todoResponse{
		baseResponse: baseResponse{
			Msg: "success",
		},
		Todo: t,
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// getTodo gets todo from store
func (ts *service) getTodo(c *fiber.Ctx) error {
	guid := c.Params("guid")

	t, err := ts.r.GetTodo(guid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	resp := todoResponse{
		baseResponse: baseResponse{
			Msg: "success",
		},
		Todo: t,
	}

	return c.JSON(resp)
}

// deleteTodo deletes todo from store
func (ts *service) deleteTodo(c *fiber.Ctx) error {
	guid := c.Params("guid")

	err := ts.r.DeleteTodo(guid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	resp := baseResponse{
		Msg: "success",
	}

	return c.Status(fiber.StatusNoContent).JSON(resp)
}

// makeDone marks todo as done
func (ts *service) makeDone(c *fiber.Ctx) error {
	guid := c.Params("guid")

	err := ts.r.MakeDone(guid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	resp := baseResponse{
		Msg: "success",
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
