package types

// Task defines the tasks common operations
type Task interface {
	IsStored() bool
}

// TaskHandler handle the tasks operations
type TaskHandler interface {
	AddTask(taskToAdd Task)
	AddTasks(tasksToAdd []Task)
	Tasks() []Task
}
