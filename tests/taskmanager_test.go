package tests

import (
	"test/taskmanager"
	"testing"
)

func Test_create(t *testing.T) {
	var all_task_data = [][]string{
		// []string{"1", "New text", "Yuriy", "Igor"},
		// []string{"asd", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
		[]string{"1", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
		[]string{"2", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
		[]string{"3", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
		[]string{"4", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
		[]string{"5", "New text", "Yuriy", "Igor", "2018-10-10 00:00:00"},
	}

	for _, single_data := range all_task_data {
		err := taskmanager.Create(single_data)
		if err != nil {
			t.Errorf("You passed: %v, and got Error: \n%s ", single_data, err)
		}
	}
}

func Test_remove(t *testing.T) {
	var tasks_id = [][]string{
		[]string{"1"},
		[]string{"2"},
		[]string{"3"},
	}

	for _, id := range tasks_id {
		err := taskmanager.Remove(id)
		if err != nil {
			t.Error(err)
		}
	}
}

func Test_list(t *testing.T) {
	tasks, err := taskmanager.List()
	if err != nil {
		t.Error(err)
		return
	}
	if len(tasks) == 0 {
		t.Error("Function returned no tasks")
	}
}
