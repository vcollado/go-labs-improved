package tasks

// TransientTask is a task in memory
type TransientTask struct {
	Task
}

// NewTransientTask creates a new transient task with its basic configuration
func NewTransientTask(name string, description string) *TransientTask {
	return &TransientTask{Task{Name: name, Description: description}}
}
