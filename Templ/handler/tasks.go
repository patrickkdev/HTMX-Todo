package handler

import (
	"htmx/templ/models"
	"htmx/templ/session"
	"htmx/templ/views"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func NewTask(c *fiber.Ctx) (models.TaskList, error) {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newTask, err := tasks.Add(models.NewTask(models.NewTaskParams{
		Title:       "Nova Tarefa",
		Description: "Descrição da tarefa",
		IsDone:      false,
	}))

	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return tasks, adaptor.HTTPHandler(templ.Handler(views.TaskItem(newTask)))(c)
}

func DeleteTask(c *fiber.Ctx) (models.TaskList, error) {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if err := tasks.Remove(id); err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return tasks, c.SendStatus(fiber.StatusAccepted)
}

func UpdateTask(c *fiber.Ctx) (models.TaskList, error) {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return tasks, c.Status(400).SendString("Invalid task ID")
	}

	title := c.FormValue("title")
	description := c.FormValue("description")
	isDone := c.FormValue("isDone") == "on"

	updatedTask := models.NewTask(models.NewTaskParams{
		ID:          id,
		Title:       title,
		Description: description,
		IsDone:      isDone,
	})

	updatedTask, err = tasks.Update(updatedTask)
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString("Failed to update task")
	}

	return tasks, adaptor.HTTPHandler(templ.Handler(views.TaskItem(updatedTask)))(c)
}
