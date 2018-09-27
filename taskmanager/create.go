package taskmanager

import (
	"fmt"
	"strconv"
	"time"
)

//Create single task
func Create(data []string) error {
	if len(data) == 0 {
		return fmt.Errorf("No data passed to create new task")
	}

	if len(data) != 5 {
		return fmt.Errorf("Wrong data passed to create task. Expected: id/text/assigner/assignee/due_date")
	}

	if _, err := strconv.Atoi(data[0]); err != nil {
		return fmt.Errorf("ID expected to be a number. Wrong data passed to create task. Expected: id/text/assigner/assignee/due_date")
	}

	var task = []string{data[0]}
	task = append(task, "true")
	task = append(task, data[1:]...)
	task = append(task, time.Now().Format("2006-01-02 15:04:05"))

	__tasks = append(__tasks, task)

	__write()
	return nil
}
