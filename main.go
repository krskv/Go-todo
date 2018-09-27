package main

import (
	"fmt"
	"os"
	"strings"
	"test/taskmanager"
)

func main() {

	if flag, data, err := get_data_from_cli(); err != nil {
		panic(err)
	} else {

		if response, err := taskmanager.Start(flag, data); err != nil {
			panic(err)
		} else {
			for _, val := range response {
				fmt.Println(val)
			}
		}
	}

	// reader := bufio.NewReader(os.Stdin)
	// if input, _ := reader.ReadString('\n'); input != "" {
	// 	return
	// }
}

//get_data_from_cli gets flag and slice of parameters
func get_data_from_cli() (string, []string, error) {
	if len(os.Args) <= 1 {
		return "", nil, fmt.Errorf("No data handled in command line. Expected flag or flag and additional data. Type -h for help")
	}

	flag := strings.Replace(os.Args[1], "-", "", -1)

	if len(os.Args) == 2 {
		return flag, nil, nil
	}

	return flag, strings.Split(strings.Replace(os.Args[2], "_", " ", -1), "/"), nil
}
