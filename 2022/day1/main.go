package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var total int
	var topThree [3]int
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			if total > topThree[0] {
				topThree[2] = topThree[1]
				topThree[1] = topThree[0]
				topThree[0] = total
			} else if total > topThree[1] {
				topThree[2] = topThree[1]
				topThree[1] = total
			} else if total > topThree[2] {
				topThree[2] = total
			}
			total = 0
		} else {
			number, _ := strconv.Atoi(text)
			total = total + number
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(topThree[0] + topThree[1] + topThree[2])
}
