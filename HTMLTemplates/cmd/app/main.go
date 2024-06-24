package main

import (
	"htmx/templates/handler"
	"htmx/templates/session"
	"time"

	"github.com/gofiber/fiber/v2"
	fsession "github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	session.Store = fsession.New(fsession.Config{
		Expiration: 24 * time.Hour,
	})

	app.Use(session.HandleSessionCookies)

	app.Static("/public", "./public")

	handler.Register(app)

	app.Listen(":3000")
}
