package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 3: Gear Ratios ---\n\n")

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
	matrix := buildMatrix(input)

	diffs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	seen := make(map[[2]int]bool)

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[row]); column++ {
			if isSymbol(matrix[row][column]) {
				for _, diff := range diffs {
					searchCoords := [2]int{row + diff[0], column + diff[1]}

					if !seen[searchCoords] && isDigit(matrix[searchCoords[0]][searchCoords[1]]) {
						result += getNumber(matrix, searchCoords, seen)
					}
				}

			}
		}
	}

	return result
}

func Part2(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	matrix := buildMatrix(input)

	diffs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	seen := make(map[[2]int]bool)

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[row]); column++ {
			if matrix[row][column] == "*" {
				numbersDetected := []int{}

				for _, diff := range diffs {
					searchCoords := [2]int{row + diff[0], column + diff[1]}

					if !seen[searchCoords] && isDigit(matrix[searchCoords[0]][searchCoords[1]]) {
						numbersDetected = append(numbersDetected, getNumber(matrix, searchCoords, seen))
					}
				}

				if len(numbersDetected) > 1 {
					sum := 1
					for _, number := range numbersDetected {
						sum *= number
					}

					result += sum
				}
			}
		}
	}

	return result
}

func buildMatrix(input string) [][]string {
	lines := strings.Split(input, "\n")
	matrix := make([][]string, len(lines))

	for row, line := range lines {
		characters := strings.Split(line, "")
		matrix[row] = make([]string, len(characters))

		for column, character := range characters {
			matrix[row][column] = character
		}
	}

	return matrix
}

func isDigit(value string) bool {
	return value >= "0" && value <= "9"
}

func isSymbol(value string) bool {
	return value != "." && !isDigit(value)
}

func getNumber(matrix [][]string, detectedCoords [2]int, seen map[[2]int]bool) int {
	row, column := detectedCoords[0], detectedCoords[1]
	for column-1 >= 0 {
		if isDigit(matrix[row][column-1]) {
			column--
		} else {
			break
		}
	}

	number := ""
	for column < len(matrix[row]) && isDigit(matrix[row][column]) {
		number += matrix[row][column]
		seen[[2]int{row, column}] = true
		column++
	}

	numberInt, err := strconv.Atoi(number)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse %v to int\n", number)
	}

	return numberInt
}
