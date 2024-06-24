package handler

import (
	"htmx/templates/session"
	"htmx/templates/views"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Render(views.Index, fiber.Map{"Title": "Minhas Tarefas", "Tasks": tasks})
}
