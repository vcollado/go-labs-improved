package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/vcollado/go-todo-list/tasks"
	"github.com/vcollado/go-todo-list/types"
)

const goMaxProcs int = 2

func main() {

	runtime.GOMAXPROCS(goMaxProcs)

	var task1, task2, task3, task4 types.Task
	var task5 types.TransientTask
	var task6 types.Task

	task1 = tasks.NewTransientTask("foo", "this is the first task")
	task2 = tasks.NewTransientTask("baz", "this is the second task")
	task3 = tasks.NewTransientTask("bas", "this is the third task")
	task4 = tasks.NewTransientTask("bar", "this is the fourth task")
	task5 = tasks.NewTransientTask("fus", "this is the fifth task")

	task6 = tasks.NewTask(task5.Name(), task5.Description())

	tasksGroup := []types.Task{task3, task4, task6}

	addTaskChannel := make(chan types.Task)
	addTasksChannel := make(chan []types.Task)

	var wg sync.WaitGroup

	var taskHandler types.TaskHandler
	taskHandler = tasks.NewTaskHandlerOverChannels(
		&wg,
		addTaskChannel,
		addTasksChannel,
	)

	addTaskChannel <- task1
	addTaskChannel <- task2
	addTasksChannel <- tasksGroup

	// wait for the goroutines to finish
	time.Sleep(time.Second)
	wg.Wait()

	println("taskHandler total tasks: ", len(taskHandler.Tasks()))
	// for _, task := range taskHandler.Tasks() {
	// 	fmt.Printf("%+v\n", task.Name())
	// }

	fmt.Printf("main finished")
}
