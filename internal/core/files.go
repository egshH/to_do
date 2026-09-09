package core

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func SaveTask() error {
	//createIfNotExist()
	jsonTask, err := json.Marshal(listOfTasks)
	if err != nil {
		return err
	}
	
	// err = os.WriteFile(fileName, jsonTask, 0644)
	// return err
	return os.WriteFile(fileName, jsonTask, 0644)
}

func LoadTask() error {
	if err := createIfNotExist(); err != nil {
		return err
	}

	jsonTask, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	if len(jsonTask) == 0 {
		return nil
	}

	tmpList := make(map[int]*task)

	err = json.Unmarshal(jsonTask, &tmpList)

	if err != nil {
		return err
	}

	listOfTasks = tmpList

	return nil
}

func createIfNotExist() error {
	_, err := os.Stat(fileName)
	if err != nil {
		_, a := os.Create(fileName)

		return a
	}
	return nil
}