package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 8: Treetop Tree House ---\n\n")

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
    size := len(strings.Split(input, "\n"))

	var treeMap [3000][3000]int
    var result [3000][3000]int

	for i, row := range strings.Split(input, "\n") {
		for j, tree := range strings.Split(row, "") {
			height, err := strconv.Atoi(tree)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse tree height: %v", err)
				os.Exit(-1)
			}
			treeMap[i][j] = height
		}
	}

    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if j-1 < 0 || treeMap[i][j] > treeMap[i][j-1]{
                result[i][j] = 1
            } else {
                if treeMap[i][j] != treeMap[i][j-1] {
                    break
                }
            }
        }
    }

    for i := 0; i < size; i++ {
        for j := size-1; j >= 0; j-- {
            if j+1 > size-1 || treeMap[i][j] > treeMap[i][j+1]{
                result[i][j] = 1
            } else {
                if treeMap[i][j] != treeMap[i][j+1] {
                    break
                }
            }
        }
    }

    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            if j-1 < 0 || treeMap[j][i] > treeMap[j-1][i]{
                result[j][i] = 1
            } else {
                if treeMap[i][j] != treeMap[j-1][i] {
                    break
                }
            }
        }
    }

    for i := 0; i < size; i++ {
        for j := size-1; j >= 0; j-- {
            if j+1 > size-1 || treeMap[j][i] > treeMap[j+1][i]{
                result[j][i] = 1
            } else {
                if treeMap[i][j] != treeMap[j+1][i] {
                    break
                }
            }
        }
    }

    sum := 0
    for row := 0; row < size; row++ {
        for column := 0; column < size; column++ {
            fmt.Fprintf(os.Stdout, "%d", result[row][column])
            if result[row][column] == 1 {
                sum += 1
            }
        }
        fmt.Fprintf(os.Stdout, "\n")
    }

	return sum 
}

func Part2(input string) int {

	return -1
}
