package types

// Task defines the tasks common operations
type Task interface {
	Name() string
	Description() string
	SetDescription(description string)
}

// TransientTask is a task in memory
type TransientTask interface {
	Task
	IsStored() bool
}

// TaskHandler handle the tasks operations
type TaskHandler interface {
	AddTask(taskToAdd Task)
	AddTasks(tasksToAdd []Task)
	EditTask(taskToEdit Task, pos int)
	Tasks() []Task
}
