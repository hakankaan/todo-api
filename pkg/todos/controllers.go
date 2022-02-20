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

// addTodo validates request with AddTodoRequest then adds todo to store
func (ts *service) addTodo(c *fiber.Ctx) error {
	var r AddTodoRequest
	if err := c.BodyParser(&r); err != nil {
		ts.l.Error("c.BodyParser", err)
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
	err := ts.pr.AddTodo(t)
	if err != nil {
		ts.l.Error("pr.AddTodo", err)
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	err = ts.rr.AddTodo(t)
	if err != nil {
		ts.l.Error("rr.AddTodo", err)
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

	t, err := ts.rr.GetTodo(guid)
	if err != nil {
		ts.l.Warn("r.GetTodo", err.Error())
		t, err = ts.pr.GetTodo(guid)
		if err != nil {
			ts.l.Error("r.GetTodo", err)
			return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
				Msg: err.Error(),
			})
		}
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

	err := ts.rr.DeleteTodo(guid)
	if err != nil {
		ts.l.Error("r.DeleteTodo", err)
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	err = ts.pr.DeleteTodo(guid)
	if err != nil {
		ts.l.Error("r.DeleteTodo", err)
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

	err := ts.rr.MakeDone(guid)
	if err != nil {
		ts.l.Error("r.MakeDone", err)
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	err = ts.pr.MakeDone(guid)
	if err != nil {
		ts.l.Error("r.MakeDone", err)
		return c.Status(fiber.StatusInternalServerError).JSON(baseResponse{
			Msg: err.Error(),
		})
	}

	resp := baseResponse{
		Msg: "success",
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
