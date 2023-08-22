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
    input = strings.TrimSuffix(input, "\n")
    size := len(strings.Split(input, "\n"))

    treeMatrix := make([][]int, size)
    visibilityMatrix := make([][]int, size)

	for i, row := range strings.Split(input, "\n") {
        treeMatrix[i] = make([]int, size)
        visibilityMatrix[i] = make([]int, size)

		for j, tree := range strings.Split(row, "") {
			height, err := strconv.Atoi(tree)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse tree height: %v", err)
				os.Exit(-1)
			}
			treeMatrix[i][j] = height
		}
	}

    for i := 0; i < size; i++ {
        rightDirectionBound := -1 
        leftDirectionBound := -1 
        downDirectionBound := -1 
        upDirectionBound := -1 

        for j := 0; j < size; j++ {
            if treeMatrix[i][j] > rightDirectionBound {
                visibilityMatrix[i][j] = 1               
                rightDirectionBound = treeMatrix[i][j]
            }

            if treeMatrix[i][size-1-j] > leftDirectionBound {
                visibilityMatrix[i][size-1-j] = 1               
                leftDirectionBound = treeMatrix[i][size-1-j]
            }

            if treeMatrix[j][i] > downDirectionBound {
                visibilityMatrix[j][i] = 1               
                downDirectionBound = treeMatrix[j][i]
            }

            if treeMatrix[size-1-j][i] > upDirectionBound {
                visibilityMatrix[size-1-j][i] = 1               
                upDirectionBound = treeMatrix[size-1-j][i]
            }
        }
    }

    sum := 0
    for row := 0; row < size; row++ {
        for column := 0; column < size; column++ {
            if visibilityMatrix[row][column] == 1 {
                sum += 1
            }
        }
    }

	return sum 
}

func Part2(input string) int {

	return -1
}
