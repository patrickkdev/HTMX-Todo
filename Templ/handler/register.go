package handler

import (
	"htmx/templ/models"
	"htmx/templ/session"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Register registers routes with handlers and ensures that the session is updated.
func Register(app *fiber.App) {
	app.Get("/", Index)

	delayMiddleWare := func(c *fiber.Ctx) error {
		time.Sleep(200 * time.Millisecond)
		return c.Next()
	}

	tasksGroup := app.Group("/tasks", delayMiddleWare)
	tasksGroup.Put("/", WithUpdateSession(NewTask))
	tasksGroup.Post("/:id", WithUpdateSession(UpdateTask))
	tasksGroup.Delete("/:id", WithUpdateSession(DeleteTask))
}

// WithUpdateSession is a middleware that calls the handler and then updates the session with the new tasks.
func WithUpdateSession(handler func(*fiber.Ctx) (models.TaskList, error)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tasks, err := handler(c)
		if err != nil {
			return err
		}

		if err := session.Save(c, tasks); err != nil {
			return err
		}

		return err
	}
}
