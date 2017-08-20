package tasks

// TransientTask is a task in memory
type TransientTask struct {
	Task
}

// NewTransientTask creates a new transient task with its basic configuration
func NewTransientTask(name string, description string) *TransientTask {
	return &TransientTask{Task{name, description, false}}
}

// IsStored checks if the task is stored
func (transientTask TransientTask) IsStored() bool {
	return transientTask.stored
}
