package handler

import (
	"htmx/templ/session"
	"htmx/templ/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func Index(c *fiber.Ctx) error {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return adaptor.HTTPHandler(templ.Handler(views.Index("Minhas Tarefas", tasks)))(c)
}
