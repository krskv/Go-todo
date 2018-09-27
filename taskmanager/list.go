package taskmanager

import (
	"fmt"
	"strings"
)

//List returns all list of tasks
func List() ([]string, error) {
	if len(__tasks) == 0 {
		return nil, fmt.Errorf("No tasks in list")
	}

	var res []string

	for _, el := range __tasks {
		res = append(res, strings.Replace(strings.Join(el, "/"), " ", "_", -1))
	}

	return res, nil
}
