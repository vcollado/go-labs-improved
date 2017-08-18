package tasks

// Task keeps the basic information
type Task struct {
	Name, Description string
}

// IsStored checks if the task is stored
func (t *Task) IsStored() bool {
	return true
}
