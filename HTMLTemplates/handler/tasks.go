package handler

import (
	"htmx/templates/models"
	"htmx/templates/session"

	"github.com/gofiber/fiber/v2"
)

func NewTask(c *fiber.Ctx) (models.TaskList, error) {
	tasks, err := session.GetTasks(c)
	if err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	newTaskID := tasks.Add(models.NewTask(models.NewTaskParams{
		Title:       "Nova Tarefa",
		Description: "Descrição da tarefa",
		IsDone:      false,
	}))

	return tasks, tasks.RenderItem(c, newTaskID)
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

	if err := tasks.Update(updatedTask); err != nil {
		return tasks, c.Status(fiber.StatusInternalServerError).SendString("Failed to update task")
	}

	return tasks, tasks.RenderItem(c, id)
}
