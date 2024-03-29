package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Print("--- Day 6: Tuning Trouble ---\n\n")

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

func hasDuplicates(slice []string) bool {
    for i, item := range slice {
        for j, itemToCompare := range slice {
            if item == itemToCompare && i != j {
                return true
            }
        }
    }

    return false 
}

func Part1(input string) int {
    letters := strings.Split(input, "")
    markers := make([]string, 4)

    for i := 0; i < len(letters); i ++ {
        markers[i % 4] = letters[i]  

        if !hasDuplicates(markers) && i >= 3 {
            return i + 1
        }
    }

    return -1 
}

func Part2(input string) int {
    letters := strings.Split(input, "")
    markers := make([]string, 14)

    for i := 0; i < len(letters); i ++ {
        markers[i % 14] = letters[i]  

        if !hasDuplicates(markers) && i >= 13 {
            return i + 1
        }
    }

    return -1 
}
