package core

import (
	"fmt"
	"slices"
)

type task struct {
	title     string
	completed bool
}

var listOfTasks map[int]*task = make(map[int]*task)

func newTask(title string) error {
	if title == "" {
		return fmt.Errorf("title must be filled")
	}

	t := &task{
		title:     title,
		completed: false,
	}

	id := len(listOfTasks) + 1

	listOfTasks[id] = t

	return nil
}

func changeTask(id int, title string) error {
	if t, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("task with number: %d is not exist", id)
	} else {
		if title == "" {
			return fmt.Errorf("title is not provided")
		}
		t.title = title
		return nil
	}
}

func reindexTask() {
	tmp := make(map[int]*task)

	i := 1

	for _, t := range listOfTasks {
		tmp[i] = t
		i++
	}

	listOfTasks = tmp
}

func deleteTask(id int) error {
	if _, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("task with number: %d is not exist", id)
	} else {
		delete(listOfTasks, id)
		reindexTask()
		return nil
	}
}

func comleteTask(id int) error {
	if t, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("task with number: %d is not exist", id)
	} else {
		t.completed = true
		return nil
	}
}

func printTasks() {
	keys := make([]int, 0, len(listOfTasks))

	for k := range listOfTasks {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for _, k := range keys {
		t := listOfTasks[k]
		checkbox := ""

		switch t.completed {
		case true:
			checkbox = "[X]"
		case false:
			checkbox = fmt.Sprintf("[%d]", k)
		}
		fmt.Printf("%s - %s\n", checkbox, t.title)
	}
}
