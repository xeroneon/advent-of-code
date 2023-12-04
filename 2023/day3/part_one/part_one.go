package partone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve(file *os.File) {
	directions := [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	var board []string
	var total int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		board = append(board, strings.TrimSpace(text))
	}

	runeBoard := make([][]rune, len(board))
	for i, row := range board {
		runeBoard[i] = []rune(row)
	}

	for rowIdx, row := range runeBoard {
		var valid bool
		var partNumber string
		for colIdx, char := range row {
			if !unicode.IsDigit(char) {
				if valid && len(partNumber) > 0 {
					fmt.Println(partNumber)
					value, err := strconv.Atoi(partNumber)
					if err != nil {
						log.Fatal(err)
					}
					total = total + value
				}
				partNumber = ""
				valid = false
				continue
			}

			partNumber = partNumber + string(char)

			for _, dir := range directions {
				newRowIdx := rowIdx + dir[0]
				newColIdx := colIdx + dir[1]
				if !isInBounds(newRowIdx, newColIdx, runeBoard) {
					continue
				}
				cell := runeBoard[newRowIdx][newColIdx]
				if !unicode.IsDigit(cell) && rune(cell) != '.' {
					valid = true
				}
			}

			if colIdx+1 == len(row) {
				if valid && len(partNumber) > 0 {
					value, err := strconv.Atoi(partNumber)
					if err != nil {
						log.Fatal(err)
					}
					total = total + value
				}
			}
		}
	}

	fmt.Println(total)
}

func isInBounds(row, col int, matrix [][]rune) bool {
	return row >= 0 && row < len(matrix[0]) && col >= 0 && col < len(matrix)
}
