package core

import (
	"encoding/json"
	"os"
)

const fileName = "tasks.json"

func SaveTask() error {
	if !isInit() {
		return nil
	}

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
	if !isInit() {
		return nil
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

func _init() error {
	_, a := os.Create(fileName)

	return a
}

func isInit() bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func _dismiss() error {
	return os.Remove(fileName)
}