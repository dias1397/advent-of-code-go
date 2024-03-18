package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("--- Day 4: Scratchcards ---\n\n")

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

func Part1(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	input = strings.ReplaceAll(input, "  ", " ")

	regex := regexp.MustCompile(`.*?: `)
	input = regex.ReplaceAllString(input, "")

	for _, game := range strings.Split(input, "\n") {
		winningNumbers := strings.Split(game, " |")[0]
		chosenNumbers := strings.Split(game, " |")[1] + " "
		gameResult := 0

		for _, winningNumber := range strings.Split(winningNumbers, " ") {
			if strings.Contains(chosenNumbers, " "+winningNumber+" ") {
				if gameResult == 0 {
					gameResult = 1
				} else {
					gameResult *= 2
				}
			}
		}

		result += gameResult
	}

	return result
}

func Part2(input string) int {

	return -1
}
