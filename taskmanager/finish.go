package taskmanager

import (
	"fmt"
)

//Finish removes single task from list by ID
func Finish(data []string) error {
	if len(data) == 0 {
		return fmt.Errorf("No ID passed to finish task")
	}

	for i, val := range __tasks {
		if val[0] == data[0] {
			__tasks[i][1] = "false"
			__write()
			return nil
		}
	}
	return fmt.Errorf("No task with such ID to remove")
}
