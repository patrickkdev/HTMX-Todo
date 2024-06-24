package main

import (
	"htmx/templ/handler"
	"htmx/templ/session"
	"time"

	"github.com/gofiber/fiber/v2"
	fsession "github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	app := fiber.New(fiber.Config{})

	session.Store = fsession.New(fsession.Config{
		Expiration: 24 * time.Hour,
	})

	app.Use(session.HandleSessionCookies)

	app.Static("/public", "./public")

	handler.Register(app)

	app.Listen(":3000")
}
