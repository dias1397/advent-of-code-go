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

func Part1(input string) int {
	input = strings.TrimSuffix(input, "\n")
	result := 0

	matrix := make([][]string, len(strings.Split(input, "\n")))
	for i, line := range strings.Split(input, "\n") {
		matrix[i] = make([]string, len(strings.Split(line, "")))
		for j, value := range strings.Split(line, "") {
			matrix[i][j] = value
		}
	}

	number := ""
	startCoord := -1
	endCoord := -1

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if isDigit(matrix[x][y]) && number == "" {
				number += matrix[x][y]
				startCoord = y - 1
				endCoord = y + 1
				continue
			}

			if isDigit(matrix[x][y]) && number != "" {
				number += matrix[x][y]
				endCoord = y + 1
			}

			if (!isDigit(matrix[x][y]) || y == len(matrix[x])-1) && number != "" {
				symbol := ""
				detectedSymbol := false
				if startCoord >= 0 && matrix[x][startCoord] != "." {
					if !isDigit(matrix[x][startCoord]) {
						detectedSymbol = true
						symbol += matrix[x][startCoord]
					}
				}
				if endCoord < len(matrix[x]) && matrix[x][endCoord] != "." {
					if !isDigit(matrix[x][endCoord]) {
						detectedSymbol = true
						symbol += matrix[x][endCoord]
					}
				}
				for i := startCoord; i <= endCoord; i++ {
					if i >= 0 && i < len(matrix[x]) {
						if x-1 >= 0 && matrix[x-1][i] != "." {
							if !isDigit(matrix[x-1][i]) {
								detectedSymbol = true
								symbol += matrix[x-1][i]
							}
						}
						if x+1 < len(matrix[x]) && matrix[x+1][i] != "." {
							if !isDigit(matrix[x+1][i]) {
								detectedSymbol = true
								symbol += matrix[x+1][i]
							}
						}
					}
				}

				if detectedSymbol {
					numberInt, err := strconv.Atoi(number)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Not a number %v", number)
					}

					result += numberInt
					fmt.Fprintf(os.Stdout, "Part number: %d -> %v\n", numberInt, symbol)
					if numberInt == 392 {
						fmt.Fprintf(os.Stdout, "%v\n%v\n%v", matrix[x-1], matrix[x], matrix[x+1])
					}
				}

				number = ""
				startCoord = -1
				endCoord = -1
			}

		}
	}

	return result
}

func detectSymbol(startY, endY int, currentLine []string) bool {
	if startY <= 0 && currentLine[startY] != "." {
		return true
	}

	if endY >= len(currentLine) && currentLine[endY] != "." {
		return true
	}

	// TODO: validate if on top line or bottom line there is a symbol

	// TODO: add number to sum

	return false
}

func isDigit(value string) bool {
	_, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return true
}

func Part2(input string) int {

	return -1
}
