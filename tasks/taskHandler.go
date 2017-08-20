package tasks

import (
	"fmt"
	"sync"
	"time"

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

// NewTaskHandlerOverChannels todo...
func NewTaskHandlerOverChannels(
	wg *sync.WaitGroup,
	addTaskCh chan types.Task,
	addTasksCh chan []types.Task,
) *TaskHandler {
	var taskHandler = NewTaskHandler()

	go func(addTaskCh chan types.Task) {
		for task := range addTaskCh {
			wg.Add(1)
			go func(task types.Task) {
				taskHandler.AddTask(task)
				wg.Done()
			}(task)
		}
	}(addTaskCh)

	go func(addTasksCh chan []types.Task) {
		for tasks := range addTasksCh {
			// V1
			wg.Add(1)
			taskHandler.AddTasks(tasks)
			wg.Done()
			// V2 todo...
			// for _, task := range tasks {
			// 	addTaskCh <- task
			// }
		}
	}(addTasksCh)

	return taskHandler
}

// AddTask to the task list
func (taskHandler *TaskHandler) AddTask(taskToAdd types.Task) {
	if len(taskHandler.tasks) == 0 {
		taskHandler.tasks = make([]types.Task, 0)
	}
	fmt.Println(time.Now().String())
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
