package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 02: Cube Conundrum ---\n\n")

	if len(os.Args) > 1 {
		input, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to read file named %s", os.Args[1])
			os.Exit(-1)
		}

		fmt.Fprintf(os.Stdout, "Part1 solution: %d\n", Part1(string(input)))
		fmt.Fprintf(os.Stdout, "Part2 solution: %d\n", Part2(string(input)))
	} else {
		fmt.Fprintf(os.Stderr, "Input file not provided")
	}
}

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	result := 0

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	gameIsPossible := true

	for index, game := range strings.Split(input, "\n") {
		game = strings.Split(game, ":")[1]

		for _, set := range strings.Split(game, ";") {
			for _, cubes := range strings.Split(set, ",") {
				color := strings.Fields(cubes)[1]
				number, err := strconv.Atoi(strings.Fields(cubes)[0])
				if err != nil {
					os.Exit(-1)
				}

				if limits[color] < number {
					gameIsPossible = false
				}
			}
		}

		if gameIsPossible {
			result += index + 1
		}
		gameIsPossible = true
	}

	return result
}

func Part2(input string) int {
	input = strings.TrimSuffix(input, "\n")
	result := 0

	cubesPerColor := map[string]int{
		"red":   -1,
		"green": -1,
		"blue":  -1,
	}

	for _, game := range strings.Split(input, "\n") {
		game = strings.Split(game, ":")[1]

		for _, set := range strings.Split(game, ";") {
			for _, cubes := range strings.Split(set, ",") {
				color := strings.Fields(cubes)[1]
				number, err := strconv.Atoi(strings.Fields(cubes)[0])
				if err != nil {
					os.Exit(-1)
				}

				if cubesPerColor[color] < number {
					cubesPerColor[color] = number
				}
			}
		}

		result += (cubesPerColor["red"] * cubesPerColor["green"] * cubesPerColor["blue"])

		cubesPerColor["red"] = -1
		cubesPerColor["green"] = -1
		cubesPerColor["blue"] = -1
	}

	return result 
}
