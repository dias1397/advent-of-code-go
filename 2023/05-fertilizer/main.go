package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 5: If You Give A Seed A Fertilizer ---\n\n")

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

	partitions := strings.Split(input, "\n\n")

	seeds := make(map[int]int)
	for _, seed := range strings.Split(partitions[0], " ") {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			continue
		}

		seeds[seedInt] = seedInt
	}

	for _, mapping := range partitions[1:] {
		for key, value := range seeds {
			for _, line := range strings.Split(mapping, "\n")[1:] {
				var destination int
				var source int
				var length int

				if _, err := fmt.Sscanf(line, "%d %d %d", &destination, &source, &length); err != nil {
					fmt.Fprintf(os.Stderr, "Error while parsing mapping values: %v", err)
				}

				if value >= source && value < source+length {
					seeds[key] = destination + (value - source)
				}
			}
		}
	}

	result := math.MaxInt
	for _, value := range seeds {
		if value < result {
			result = value
		}
	}

	return result
}

func Part2(input string) int {

	return -1
}
