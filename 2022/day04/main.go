package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Print("Advent of code 2022 - Day 4\n\n")

    if len(os.Args) > 1 {
        input := readInput(os.Args[1])
        fmt.Fprintf(os.Stdout, "Part 1 solution: %d\n", Part1(input))

        fmt.Fprintf(os.Stdout, "Part 2 solution: %d\n", Part2(input))
    } else {
        fmt.Fprintf(os.Stderr, "Input file not provided") 
    }
}

func readInput(filename string) string {
    fileContentByteSlice, err := os.ReadFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to read file named %s\n", filename)
        os.Exit(-1)
    }

    return string(fileContentByteSlice)
}

func Part1(input string) int {
    result := 0

    for _, line := range strings.Fields(input) {
        sectionsByElf := strings.Split(line, ",")
        bounds := []int{}
        
        for _, sections := range sectionsByElf {
            for _, section := range strings.Split(sections, "-") {
                temp, err := strconv.Atoi(section)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Error in %s conversion to int", section)
                }

                bounds = append(bounds, temp)
            }
        }

        if (bounds[0] >= bounds[2] && bounds[1] <= bounds[3]) || 
            (bounds[2] >= bounds[0] && bounds[3] <= bounds[1]) {
            result += 1    
        }
    }

    return result
}

func Part2(input string) int {
    result := 0

    for _, line := range strings.Fields(input) {
        sectionsByElf := strings.Split(line, ",")
        bounds := []int{}
        
        for _, sections := range sectionsByElf {
            for _, section := range strings.Split(sections, "-") {
                temp, err := strconv.Atoi(section)
                if err != nil {
                    fmt.Fprintf(os.Stderr, "Error in %s conversion to int", section)
                }

                bounds = append(bounds, temp)
            }
        }

        if (bounds[0] >= bounds[2] && bounds[0] <= bounds[3]) || 
            (bounds[1] >= bounds[2] && bounds[1] <= bounds[3]) || 
            (bounds[2] >= bounds[0] && bounds[2] <= bounds[1]) || 
            (bounds[3] >= bounds[0] && bounds[3] <= bounds[1]) {
            result += 1    
        }
    }

    return result
}
