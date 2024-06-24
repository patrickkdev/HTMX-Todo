package models

import (
	"errors"
	"math/rand"
	"sync"
)

type Task struct {
	ID          int
	Title       string
	Description string
	IsDone      bool
	mu          *sync.Mutex
}

type NewTaskParams struct {
	ID          int
	Title       string
	Description string
	IsDone      bool
}

// No need to pass ID if it's a new task. Only pass ID if it's an update to an existing task
func NewTask(params NewTaskParams) Task {
	n := Task{
		ID:          params.ID,
		Title:       params.Title,
		Description: params.Description,
		IsDone:      params.IsDone,
		mu:          &sync.Mutex{},
	}

	if n.ID == 0 {
		n.ID = rand.Int()
	}

	return n
}

// func (t task) Render(c *fiber.Ctx) error {
// return c.Render(views.TasksItem, t)
// }

type TaskList []Task

func NewTaskList() TaskList {
	newTaskList := TaskList{}

	newTaskList = append(newTaskList, Task{
		ID:          1,
		Title:       "Tarefa 1",
		Description: "Descrição da tarefa 1",
		IsDone:      false,
	})

	return newTaskList
}

func (ts TaskList) GetAll() TaskList {
	return ts
}

func (ts TaskList) Get(id int) (Task, error) {
	for _, task := range ts {
		if task.ID == id {
			return task, nil
		}
	}

	return Task{}, errors.New("task not found")
}

func (ts *TaskList) Add(task Task) (Task, error) {
	*ts = append(*ts, task)

	return ts.Get(task.ID)
}

func (ts *TaskList) Update(task Task) (Task, error) {
	for i, t := range *ts {
		if t.ID == task.ID {
			(*ts)[i] = task
			return ts.Get(task.ID)
		}
	}

	return Task{}, errors.New("task not found")
}

func (ts *TaskList) Remove(id int) error {
	for i, task := range *ts {
		if task.ID == id {
			*ts = append((*ts)[:i], (*ts)[i+1:]...)
			return nil
		}
	}

	return errors.New("task not found")
}

func (ts TaskList) ToFiltered(f func(Task) bool) TaskList {
	filtered := TaskList{}

	for _, task := range ts {
		if f(task) {
			filtered = append(filtered, task)
		}
	}

	return filtered
}

func (ts *TaskList) Filter(f func(Task) bool) {
	*ts = ts.ToFiltered(f)
}

// func (ts TaskList) RenderAll(c *fiber.Ctx, data fiber.Map) error {
// 	data["Tasks"] = ts

// 	for k, v := range data {
// 		fmt.Println(k, v)
// 	}

// 	return c.Render(views.TasksContainer, data)
// }

// func (ts TaskList) RenderItem(c *fiber.Ctx, id int) error {
// 	item, err := ts.Get(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).SendString(err.Error())
// 	}

// 	return item.Render(c)
// }
