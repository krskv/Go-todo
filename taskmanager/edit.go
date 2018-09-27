package taskmanager

import (
	"fmt"
)

//Remove removes single task from list by ID
func Edit(data []string) error {
	if len(data) <= 0 {
		return fmt.Errorf("No data passed to edit task")
	}

	if len(data) != 2 {
		return fmt.Errorf("Wrong data passed to create task. Expected: id;new text")
	}

	for i, val := range __tasks {
		if val[0] == data[0] {
			__tasks[i][2] = data[1]
			__write()
			return nil
		}
	}

	return fmt.Errorf("No task with such ID to remove")
}
