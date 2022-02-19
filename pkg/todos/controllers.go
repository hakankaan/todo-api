package todos

import "github.com/gofiber/fiber/v2"

func (ts *service) ping(c *fiber.Ctx) error {
	return c.SendString("success")
}
