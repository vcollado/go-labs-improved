package main

import (
	tasks "github.com/vcollado/go-todo-list/tasks"
	types "github.com/vcollado/go-todo-list/types"
)

func main() {

	var task types.Task

	task = tasks.NewTransientTask("name 2", "description 2")

	println(task.IsStored())
}
