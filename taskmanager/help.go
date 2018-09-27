package taskmanager

//Help returns list of all flags
func Help() ([]string, error) {
	response := []string{"-c <id/text/assigner/assignee/due_date> - creates single task", "-l - lists all tasks", "-r <id> - removes task by ID", "-e <id/new_text> - edits task text", "-f <id> - makes task done"}
	return response, nil
}
