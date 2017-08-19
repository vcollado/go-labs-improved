package main

import (
	"github.com/vcollado/go-todo-list/tasks"
	"github.com/vcollado/go-todo-list/types"
)

func main() {

	var task1, task2, task3, defaultTask types.Task

	task1 = tasks.NewTransientTask("this is the first task", "foo")
	task2 = tasks.NewTransientTask("this is the second task", "baz")
	task3 = tasks.NewTransientTask("this is the third task", "bas")
	defaultTask = tasks.NewTransientTask("this is the default task", "fus")

	var defaultTaskGroup = []types.Task{defaultTask}
	var taskGroup1 = []types.Task{task2, task3}

	var taskHandler types.TaskHandler

	taskHandler = tasks.NewTaskHandler(defaultTaskGroup...)
	taskHandler.AddTask(task1)
	taskHandler.AddTasks(taskGroup1)

	for _, task := range taskHandler.Tasks() {
		println(task.IsStored())
	}

}
