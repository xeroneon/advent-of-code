package main

import (
	parttwo "day3/part_two"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No file provided")
		return
	}

	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	defer file.Close()

	parttwo.Solve(file)
	// gpt.Solve()
}
