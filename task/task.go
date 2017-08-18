package task

const newTasksAreDoneByDefault = false

// Task is the basic type of the application
type Task struct {
	name        string
	description string
	done        bool
}

// CreateTask given a name and a description
func CreateTask(name string, description string) *Task {
	return &Task{
		name:        name,
		description: description,
		done:        newTasksAreDoneByDefault,
	}
}

func (t Task) setDone() {
	t.done = true
}

func (t Task) setNotDone() {
	t.done = false
}
