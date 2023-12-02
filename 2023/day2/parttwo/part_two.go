package parttwo

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func PartTwo(file *os.File) {
	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, ": ")

		rounds := strings.Split(parts[1], "; ")
		re := regexp.MustCompile(`(\d+) (\w+)`)

		var redMin int
		var blueMin int
		var greenMin int

		for _, round := range rounds {
			matches := re.FindAllStringSubmatch(round, -1)
			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				switch color {
				case "red":
					if count > redMin || redMin == 0 {
						redMin = count
					}
				case "blue":
					if count > blueMin || blueMin == 0 {
						blueMin = count
					}
				case "green":
					if count > greenMin || greenMin == 0 {
						greenMin = count
					}
				}
			}
		}

		total = total + (redMin * blueMin * greenMin)
	}

	fmt.Println(total)
}
