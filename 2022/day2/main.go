package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	String string
	Value  int
	Index  int
}

func main() {
	var total int
	numbers := []Number{
		{String: "one", Value: 1},
		{String: "two", Value: 2},
		{String: "three", Value: 3},
		{String: "four", Value: 4},
		{String: "five", Value: 5},
		{String: "six", Value: 6},
		{String: "seven", Value: 7},
		{String: "eight", Value: 8},
		{String: "nine", Value: 9},
	}
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first rune
		var last rune
		var firstString Number
		var lastString Number
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		for _, num := range numbers {
			startIndex := strings.Index(text, num.String)
			lastIndex := strings.LastIndex(text, num.String)
			if startIndex == -1 {
				continue
			}

			if firstString.Value == 0 {
				num.Index = startIndex
				firstString = num
			}
			if startIndex <= firstString.Index {
				num.Index = startIndex
				firstString = num
			}

			if lastIndex >= lastString.Index {
				num.Index = lastIndex
				lastString = num
			}
		}

		for _, r := range text {
			if isDigit(r) {
				first = r
				break
			}
		}

		reversed := []rune(text)

		for i := len(reversed) - 1; i >= 0; i-- {
			if isDigit(reversed[i]) {
				last = reversed[i]
				break
			}
		}

		firstIndex := strings.Index(text, string(first))
		lastIndex := strings.LastIndex(text, string(last))

		if firstIndex > firstString.Index {
			if firstString.Value != 0 {
				first = rune('0' + firstString.Value)
			}
		}

		if lastIndex < lastString.Index {
			if lastString.Value != 0 {
				last = rune('0' + lastString.Value)
			}
		}

		combinedDigit := string(first) + string(last)
		i, err := strconv.Atoi(combinedDigit)
		if err != nil {
			log.Fatal(err)
		}

		total = total + i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(total)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
