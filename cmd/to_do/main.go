package main

import (
	"fmt"
	"os"
	"to_do/internal/core"
)

func main() {
	args := os.Args
	err := core.Parse(args)
	if err != nil {
		fmt.Print(err)
	}
}
