package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 1: Trebuchet?! ---\n\n")

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

	firstDigit := -1
	secondDigit := -1

	for _, line := range strings.Split(input, "\n") {
		for _, value := range strings.Split(line, "") {
			temp, err := strconv.Atoi(value)
			if err != nil {
				continue
			}

			if firstDigit == -1 {
				firstDigit = temp
			}
			secondDigit = temp
		}
		result += ((firstDigit * 10) + secondDigit)

		firstDigit = -1
		secondDigit = -1
	}

	return result
}

func Part2(input string) int {
    digitsAsStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    for index, word := range digitsAsStrings{
        input = strings.ReplaceAll(input, word, word + strconv.Itoa(index + 1) + word) 
    }

	return Part1(input) 
}
