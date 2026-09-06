package core

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func saveTask() error {
	jsonTask, err := json.Marshal(listOfTasks)
	if err != nil {
		return err
	}
	
	err = os.WriteFile(fileName, jsonTask, 0644)

	return err
}

// func loadTask() error {
// 	return nil
// }
