package partone

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Blue  int
	Green int
	Red   int
}

const (
	RedMax   = 12
	BlueMax  = 14
	GreenMax = 13
)

func PartOne(file *os.File) {
	var validGames []Game
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		parts := strings.Split(text, ": ")
		gameNumber, _ := strconv.Atoi(strings.Split(parts[0], " ")[1])

		rounds := strings.Split(parts[1], "; ")
		re := regexp.MustCompile(`(\d+) (\w+)`)

		game := Game{Id: gameNumber}

		for i, round := range rounds {
			matches := re.FindAllStringSubmatch(round, -1)
			var red int
			var blue int
			var green int
			for _, match := range matches {
				count, _ := strconv.Atoi(match[1])
				color := match[2]
				switch color {
				case "red":
					red = count
				case "blue":
					blue = count
				case "green":
					green = count
				}
			}
			if red > RedMax || blue > BlueMax || green > GreenMax {
				break
			}
			game.Red = game.Red + red
			game.Blue = game.Blue + blue
			game.Green = game.Green + green
			if i+1 == len(rounds) {
				validGames = append(validGames, game)
			}
		}
	}

	var total int

	for _, game := range validGames {
		total = total + game.Id
	}
	fmt.Println(total)
}
