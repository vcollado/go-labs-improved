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

	var task0 types.Task = tasks.NewTransientTask("foo", "this is the first task")
	var task1 types.Task = tasks.NewTransientTask("baz", "this is the second task")
	var task2 types.Task = tasks.NewTransientTask("bas", "this is the third task")
	var task3 types.Task = tasks.NewTransientTask("bar", "this is the fourth task")
	// var task5 types.TransientTask = tasks.NewTransientTask("fus", "this is the fifth task")

	// var task6 types.Task
	// task6 = tasks.NewTask(task5.Name(), task5.Description())

	tasksGroup := []types.Task{task2, task3}

	addTaskChannel := make(chan types.Task)
	addTasksChannel := make(chan []types.Task)

	var wg sync.WaitGroup

	var taskHandler types.TaskHandler
	taskHandler = tasks.NewTaskHandlerOverChannels(
		&wg,
		addTaskChannel,
		addTasksChannel,
	)

	addTaskChannel <- task0
	addTaskChannel <- task1
	addTasksChannel <- tasksGroup

	time.Sleep(time.Second)
	wg.Wait()

	println("taskHandler total tasks: ", len(taskHandler.Tasks()))

	task2.SetDescription("edited description")
	taskHandler.EditTask(task2, 2)
	// fmt.Printf("%+v\n", taskHandler.Tasks()[3])

	for _, task := range taskHandler.Tasks() {
		fmt.Printf("%+v\n", task)
	}

	fmt.Printf("main finished")
}
