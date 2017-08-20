package tasks

// Task keeps the basic information
type Task struct {
	name, description string
	stored            bool
}

// NewTask creates a new task with its basic configuration
func NewTask(name string, description string) *Task {
	return &Task{name, description, false}
}

// Name of the task
func (t Task) Name() string {
	return t.name
}

// Description of the task
func (t Task) Description() string {
	return t.description
}
