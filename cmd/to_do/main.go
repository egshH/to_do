package main

import (
	"fmt"
	"os"
	"to_do/internal/core"
)

func main() {
	err := core.LoadTask()
	if err != nil {
		panic(err)
	}

	args := os.Args
	err = core.Parse(args)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	err = core.SaveTask()
	if err != nil {
		panic(err)
	}
}
