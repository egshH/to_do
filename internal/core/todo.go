package core

import "fmt"

type task struct {
	title     string
	completed bool
}

var listOfTasks map[int]*task = make(map[int]*task)

func NewTask(title string) error {
	if title == "" {
		return fmt.Errorf("описание должно быть заполнено!")
	}

	t := &task{ 
		title:     title,
		completed: false,
	}

	id := len(listOfTasks) + 1

	listOfTasks[id] = t

	return nil
}

func ChangeTask(id int, title string) error {
	if t, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("задачи под номером: %d не существует", id)
	} else {
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

func DeleteTask(id int) error {
	if _, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("задачи под номером: %d не существует", id)
	} else {
		delete(listOfTasks, id)
		reindexTask()
		return nil
	}
}

func ComleteTask(id int) error {
	if t, ok := listOfTasks[id]; !ok {
		return fmt.Errorf("задачи под номером: %d не существует", id)
	} else {
		t.completed = true
		return nil
	}
}
