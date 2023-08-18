package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("--- Day 7: No Space Left On Device ---\n\n")

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
	memory := make(map[string]int)

    mapMemory(input, memory)

	result := 0
	for _, totalSize := range memory {
		if totalSize <= 100000 {
			result = result + totalSize
		}
	}

	return result
}

func Part2(input string) int {
    memory := make(map[string]int)

    mapMemory(input, memory)

    freeSpace:= 70000000 - memory["/"]
    toBeFree := 30000000 - freeSpace

    result := 70000000 
    for _, totalSize := range memory{
        if totalSize >= toBeFree && totalSize <= result {
            result = totalSize 
        } 
    }

	return result
}

func mapMemory(input string, memory map[string]int) {
	var dirDepth []string

	for _, instruction := range strings.Split(input, "\n") {
		options := strings.Split(instruction, " ")

		if options[0] == "$" {
			if options[1] == "cd" {
				if options[2] == ".." {
					dirDepth = dirDepth[:len(dirDepth)-1]
				} else {
                    var newDir string
                    if len(dirDepth) > 0 {
                        newDir = dirDepth[len(dirDepth)-1] + "/" + options[2]
                    } else {
                        newDir = options[2]
                    }

					dirDepth = append(dirDepth, newDir)
                    if memory[newDir] != 0 {
                        fmt.Fprintf(os.Stdout, "%s[%d]\n", newDir, memory[newDir])
                    }
				}
			}
		} else {
			size, err := strconv.Atoi(options[0])
			if err != nil {
				continue
			}
			for _, dir := range dirDepth {
				val := memory[dir]
				memory[dir] = val + size
			}
		}
	}
}
