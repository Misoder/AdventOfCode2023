package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Misoder/AdventOfCode2023/tools"
)

type Set struct {
	Red   uint64
	Green uint64
	Blue  uint64
}

type Game struct {
	Id   int
	Sets []Set
}

func (g *Game) possible(maxRed uint64, maxGreen uint64, maxBlue uint64) bool {
	for _, s := range g.Sets {
		if s.Red > maxRed || s.Green > maxGreen || s.Blue > maxBlue {
			return false
		}
	}
	return true
}

func (g *Game) minCubeSet() Set {
	var red, green, blue uint64 = 0, 0, 0

	for _, cubes := range g.Sets {
		if cubes.Red > red {
			red = cubes.Red
		}
		if cubes.Green > green {
			green = cubes.Green
		}
		if cubes.Blue > blue {
			blue = cubes.Blue
		}
	}

	return Set{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func parseLine(line string) (*Game, error) {
	splits := strings.Split(line, ":")
	gameId, err := strconv.Atoi(splits[0][5:])
	if err != nil {
		return nil, err
	}

	game := &Game{Id: gameId}
	sets := strings.Split(splits[1], ";")

	for _, set := range sets {
		set = strings.TrimSpace(set)
		cubes := strings.Split(set, ",")
		newSet := &Set{}

		for _, c := range cubes {
			c = strings.TrimSpace(c)
			t := strings.Split(c, " ")

			amount, err := strconv.ParseUint(t[0], 10, 0)
			if err != nil {
				return nil, err
			}

			switch t[1] {
			case "red":
				newSet.Red = amount
			case "green":
				newSet.Green = amount
			case "blue":
				newSet.Blue = amount
			default:
				return nil, fmt.Errorf("unexpected color %s", t[1])
			}
		}
		game.Sets = append(game.Sets, *newSet)
	}

	return game, nil
}

func main() {
	t := time.Now()
	defer fmt.Printf("Duration: %s\n", time.Now().Sub(t))

	lineCh, err := tools.ReadLines("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	var powerSum uint64 = 0
	for line := range lineCh {
		game, err := parseLine(line)
		if err != nil {
			fmt.Printf("error parsing input: %v", err)
			return
		}

		if game.possible(12, 13, 14) {
			sum = sum + game.Id
		}

		minSet := game.minCubeSet()
		powerSum = powerSum + (minSet.Red * minSet.Green * minSet.Blue)
	}

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Sum of powers: %d\n", powerSum)
}
