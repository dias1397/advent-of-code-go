package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    fmt.Print("Advent of Code 2022 - Day3\n\n")

    if len(os.Args) > 1 {
        fmt.Fprintf(os.Stdout, "Part1 solution: %d\n", Part1(os.Args[1]))

        fmt.Fprintf(os.Stdout, "Part2 solution: %d\n", Part2(os.Args[1]))
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

func Part1(filename string) int {
    fileContent := readInput(filename)
    result := 0

    for _, line := range strings.Fields(fileContent) {
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

func Part2(filename string) int {
    fileContent := readInput(filename) 
    result := 0

    lines := strings.Fields(fileContent)

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
