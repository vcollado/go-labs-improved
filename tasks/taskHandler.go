package tasks

import (
	"github.com/vcollado/go-todo-list/types"
)

// TaskHandler manage the operations over tasks
type TaskHandler struct {
	tasks []types.Task
}

// NewTaskHandler creates a new task handler with its basic configuration
func NewTaskHandler(defaultTasks ...types.Task) *TaskHandler {
	return &TaskHandler{tasks: defaultTasks}
}

// AddTask to the task list
func (taskHandler *TaskHandler) AddTask(taskToAdd types.Task) {
	if len(taskHandler.tasks) == 0 {
		taskHandler.tasks = make([]types.Task, 0)
	}

	taskHandler.tasks = append(taskHandler.tasks, taskToAdd)
}

// AddTasks to the task list
func (taskHandler *TaskHandler) AddTasks(tasksToAdd []types.Task) {
	for _, task := range tasksToAdd {
		taskHandler.AddTask(task)
	}
}

// Tasks of the tasksHandler
func (taskHandler TaskHandler) Tasks() []types.Task {
	return taskHandler.tasks
}
