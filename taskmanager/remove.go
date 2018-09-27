package taskmanager

import "fmt"

//Remove removes single task from list by ID
func Remove(data []string) error {
	if len(data) == 0 {
		return fmt.Errorf("No ID handled to remove task")
	}

	if len(__tasks) == 0 {
		return fmt.Errorf("No tasks in list")
	}

	for i, val := range __tasks {
		if val[0] == data[0] {
			var res [][]string
			res = append(res, __tasks[:i]...)
			res = append(res, __tasks[i+1:]...)
			__tasks = res
			__write()
			return nil
		}
	}

	return fmt.Errorf("No task with such ID to remove")
}
