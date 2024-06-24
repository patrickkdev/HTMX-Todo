package handler

import (
	"htmx/templates/models"
	"htmx/templates/session"

	"github.com/gofiber/fiber/v2"
)

// Register registers routes with handlers and ensures that the session is updated.
func Register(app *fiber.App) {
	app.Get("/", Index)

	app.Put("/tasks", WithUpdateSession(NewTask))
	app.Post("/tasks/:id", WithUpdateSession(UpdateTask))
	app.Delete("/tasks/:id", WithUpdateSession(DeleteTask))
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
