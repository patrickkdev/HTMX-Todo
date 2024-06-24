package session

import (
	"encoding/gob"
	"errors"
	"fmt"
	"htmx/templates/config"
	"htmx/templates/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store
var sessionCookieKey = "session_id"

func init() {
	gob.Register(models.TaskList{})
}

func HandleSessionCookies(c *fiber.Ctx) error {
	userSession, err := Store.Get(c)
	if err != nil {
		return err
	}

	sessionFromRequest := c.Cookies(sessionCookieKey)
	currentSessionID := userSession.ID()

	if sessionFromRequest != currentSessionID {
		fmt.Println("Session ID does not match", sessionFromRequest, currentSessionID)

		c.Cookie(&fiber.Cookie{
			Name:  sessionCookieKey,
			Value: userSession.ID(),
		})
	}

	return c.Next()
}

func GetTasks(c *fiber.Ctx) (models.TaskList, error) {
	session, err := Store.Get(c)
	if err != nil {
		return models.TaskList{}, err
	}

	if session == nil {
		return models.TaskList{}, errors.New("no session")
	}

	tasks, ok := session.Get(config.TasksStoreKey).(models.TaskList)
	if !ok {
		tasks = models.NewTaskList()
		session.Set(config.TasksStoreKey, tasks)
		err := session.Save()
		if err != nil {
			return models.TaskList{}, err
		}
	}

	return tasks, nil
}

func Save(c *fiber.Ctx, tasks models.TaskList) error {
	session, err := Store.Get(c)
	if err != nil {
		return err
	}

	session.Set(config.TasksStoreKey, tasks)
	return session.Save()
}
