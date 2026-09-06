package core

import (
	"fmt"
	"strconv"
)

func Parse(args []string) error {
	if len(args) == 1 {
		return fmt.Errorf("command is not provided")
	}

	switch args[1] {
	case "print":
		if len(args) != 2 {
			return fmt.Errorf("to many arguments")
		}
		printTasks()
	case "add":
		if len(args) != 3 {
			return fmt.Errorf("count of arguments must be 3")
		}
		return newTask(args[2])
	case "change":
		if len(args) != 4 {
			return fmt.Errorf("count of arguments must be 4")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("%s is not number", args[2])
		}
		return changeTask(id, args[3])
	case "delete":
		if len(args) != 3 {
			return fmt.Errorf("count of arguments must be 3")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("%s is not number", args[2])
		}
		return deleteTask(id)
	case "complete":
		if len(args) != 3 {
			return fmt.Errorf("count of arguments must be 3")
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("%s is not number", args[2])
		}
		return comleteTask(id)
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}

	return nil
}
