package taskmanager

import (
	"encoding/csv"
	"fmt"
	"os"
)

var (
	__file  *os.File
	__tasks [][]string
)

//Start is entry point to all taskmanager operation
func Start(flag string, data []string) ([]string, error) {
	var response []string

	if err := __initialize(); err != nil {
		return nil, err
	}

	if res, err := __handleFlag(flag, data); err != nil {
		return nil, err
	} else {
		response = res
		return response, nil
	}
}

//__handleFlag gets flag and due to result starts an action
func __handleFlag(flag string, data []string) ([]string, error) {

	switch flag {
	case "c":
		if err := Create(data); err != nil {
			return nil, err
		}
	case "r":
		if err := Remove(data); err != nil {
			return nil, err
		}
	case "h":
		if r, err := Help(); err != nil {
			return nil, err
		} else {
			return r, nil
		}

	case "e":
		if err := Edit(data); err != nil {
			return nil, err
		}
	case "f":
		if err := Finish(data); err != nil {
			return nil, err
		}
	case "l":
		if r, err := List(); err != nil {
			return nil, err
		} else {
			return r, nil
		}
	}

	return nil, nil
}

//__initialize opens file and gets all tasks from it
func __initialize() error {

	if nil != __file {
		return fmt.Errorf("Initialize can not be called twice")
	}

	f, err := os.OpenFile("list.csv", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	__file = f

	r := csv.NewReader(__file)
	t, err := r.ReadAll()
	if err != nil {
		return err
	}

	__tasks = t

	return nil
}

//__write writes updated tasks to file
func __write() error {
	w := csv.NewWriter(__file)

	__file.Truncate(0)
	__file.Seek(0, 0)

	for _, task := range __tasks {
		if err := w.Write(task); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}
