package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Print("Advent of Code 2022 - Day1\n\n")

    if len(os.Args) > 1 {
        input := readInput(os.Args[1])
        fmt.Fprintf(os.Stdout, "Part1 solution: %d\n", Part1(input))

        fmt.Fprintf(os.Stdout, "Part2 solution: %d\n", Part2(input))
    } else {
        fmt.Fprintf(os.Stderr, "Input file not provided")
    }
}

func readInput(filename string) string {
    fileContentByteSlice, err := os.ReadFile(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to read file named %s", filename)
        os.Exit(-1)
    }

    return string(fileContentByteSlice)
}

func Part1(input string) int {
    result := 0

    scores := map[string]int{"X": 1, "Y": 2, "Z": 3, 
                             "A": 1, "B": 2, "C": 3}

    for _, line := range strings.Split(input, "\n") {
        if line == "" {
            continue
        }

        round := strings.Fields(line) 
        result += scores[round[1]]

        switch (scores[round[1]] - scores[round[0]] + 3) % 3 {
            case 0:
                result += 3
            case 1:
                result += 6
        }
    }

    return result
}

func Part2(input string) int {
    result := 0

    scores := map[string]int{"X": 0, "Y": 3, "Z": 6, 
                             "A": 1, "B": 2, "C": 3}

    for _, line := range strings.Split(input, "\n") {
        if line == "" {
            continue
        }

        round := strings.Fields(line) 
        result += scores[round[1]]

        var temp int
        switch round[1] {
            case "X":
                temp = (scores[round[0]] - 1) % 3 
            case "Y":
                temp = scores[round[0]]
            case "Z":
                temp = (scores[round[0]] + 1) % 3
        }
        
        if temp == 0 {
            temp = 3
        }

        result += temp
    }

    return result
}
