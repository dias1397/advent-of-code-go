package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Print("--- Day 3: Rucksack Reorganization ---\n\n")

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
    result := 0

    for _, line := range strings.Fields(input) {
        prefix := strings.Split(line[:(len(line)/2)], "")
        sufix  := line[(len(line)/2):]

        for _, item := range prefix {
            if strings.Contains(sufix, item) {
                value := int(item[0])
                if value > 96 && value < 123 {
                    result += (value - 96)
                }
                if value > 64 && value < 91 {
                    result += (value - 38)
                }
                break
            }
        }
    }

    return result 
}   

func Part2(input string) int {
    result := 0

    lines := strings.Fields(input)

    for i := 0; i < len(lines); i += 3 {
        for _, item := range strings.Split(lines[i], "") {
            if strings.Contains(lines[i+1], item) && 
                    strings.Contains(lines[i+2], item) {
                value := int(item[0])
                if value > 96 && value < 123 {
                    result += (value - 96)
                }
                if value > 64 && value < 91 {
                    result += (value - 38)
                }
                break
            } 
        }
    }

    return result
}
