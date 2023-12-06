package parttwo

import (
	"bufio"
	"fmt"
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
		for colIdx, char := range row {
			var first string
			var firstCoord [][2]int
			var second string
			if char == '*' {
				for _, dir := range directions {
					newRowIdx := rowIdx + dir[0]
					newColIdx := colIdx + dir[1]

					if !isInBounds(newRowIdx, newColIdx, runeBoard) {
						continue
					}

					cell := runeBoard[newRowIdx][newColIdx]

					if unicode.IsDigit(cell) {
						if len(first) != 0 && len(second) != 0 {
							continue
						}
						if len(first) == 0 {
							// store in first variable
							var left int
							right := 1

							for newColIdx-left >= 0 && unicode.IsDigit(runeBoard[newRowIdx][newColIdx-left]) {
								first = string(runeBoard[newRowIdx][newColIdx-left]) + first
								firstCoord = append(firstCoord, [2]int{newRowIdx, newColIdx - left})
								left++
							}

							for newColIdx+right < len(row) && unicode.IsDigit(runeBoard[newRowIdx][newColIdx+right]) {
								first = first + string(runeBoard[newRowIdx][newColIdx+right])
								firstCoord = append(
									firstCoord,
									[2]int{newRowIdx, newColIdx + right},
								)
								right++
							}
							left = 0
							right = 1
						} else {
							if containsCoordinates(firstCoord, [2]int{newRowIdx, newColIdx}) {
								// the order of operations for the directions means sometimes the first number is detected as the second number
								continue
							}
							// store in second variable
							var left int
							right := 1

							for newColIdx-left >= 0 && unicode.IsDigit(runeBoard[newRowIdx][newColIdx-left]) {
								second = string(runeBoard[newRowIdx][newColIdx-left]) + second
								left++
							}

							for newColIdx+right < len(row) && unicode.IsDigit(runeBoard[newRowIdx][newColIdx+right]) {
								second = second + string(runeBoard[newRowIdx][newColIdx+right])
								right++
							}
							left = 0
							right = 1
						}
					}
				}
				if len(first) == 0 || len(second) == 0 {
					continue
				}

				if first == second {
					continue
				}
				// fmt.Println("final first: ", first)
				// fmt.Println("final second: ", second)

				value1, _ := strconv.Atoi(first)
				value2, _ := strconv.Atoi(second)

				product := value1 * value2
				total = total + product
			}
		}
	}

	fmt.Println(total)
}

func isInBounds(row, col int, matrix [][]rune) bool {
	return row >= 0 && row < len(matrix[0]) && col >= 0 && col < len(matrix)
}

func containsCoordinates(board [][2]int, coord [2]int) bool {
	for _, cell := range board {
		if cell[0] == coord[0] && cell[1] == coord[1] {
			return true
		}
	}
	return false
}
