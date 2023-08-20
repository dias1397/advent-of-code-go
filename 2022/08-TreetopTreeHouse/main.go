package main

import (
	"fmt"
	"log"
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
	var treeMap [5][5]int

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

	log.Printf("%v", treeMap)

	return -1
}

func Part2(input string) int {

	return -1
}
