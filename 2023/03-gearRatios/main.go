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

func hasSymbolAroundNumber(start, end, rowIndex int, matrix [][]string) bool {
	height := len(matrix)
	width := len(matrix[rowIndex])

	// looks for symbol on left edge of number
	if start >= 0 && isSymbol(matrix[rowIndex][start]) {
		return true
	}

	// looks for symbol on right edge of number
	if end < width && isSymbol(matrix[rowIndex][end]) {
		return true
	}

	// looks for symbol on top and bottom row of number
	for column := start; column <= end; column++ {
		if column >= 0 && column < width {
			if rowIndex-1 >= 0 && isSymbol(matrix[rowIndex-1][column]) {
				return true
			}

			if rowIndex+1 < height && isSymbol(matrix[rowIndex+1][column]) {
				return true
			}
		}
	}

	return false
}

func Part1(input string) (result int) {
	input = strings.TrimSuffix(input, "\n")
	matrix := buildMatrix(input)

	number := ""
	startPos := -1
	endPos := -1

	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[row]); column++ {
			if isDigit(matrix[row][column]) {
				if number == "" {
					startPos = column - 1
				}
				endPos = column + 1
				number += matrix[row][column]
			}

			if (!isDigit(matrix[row][column]) || column == len(matrix[row])-1) && number != "" {
				if hasSymbolAroundNumber(startPos, endPos, row, matrix) {
					numberInt, err := strconv.Atoi(number)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Not a number %v", number)
					}

					result += numberInt
				}

				number = ""
				startPos = -1
				endPos = -1
			}
		}
	}

	return result
}

func Part2(input string) int {

	return -1
}
