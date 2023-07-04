package main

import (
	"fmt"
	"os"
	"strings"
)

type stack struct {
    elements []rune
}

func (s *stack) push(r rune) {
    s.elements = append(s.elements, r)
}

func (s *stack) pop() (r rune) {
    r = s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return 
}

func (s stack) top() rune {
    return s.elements[len(s.elements)-1]
}

func main() {
    fmt.Print("Advent of code 2022 - Day 5\n\n")

    if len(os.Args) > 1 {
        fmt.Fprintf(os.Stdout, "Part 1 solution: %s\n", Part1(os.Args[1]))

        fmt.Fprintf(os.Stdout, "Part 2 solution: %d\n", Part2(os.Args[1]))
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

func Part1(filename string) string {
    fileContent := readInput(filename)
    parts := strings.Split(fileContent, "\n\n")

    lines := strings.Split(parts[0], "\n")
    lines = lines[:len(lines)-1]

    // TODO remove hardcoded number of stacks
    stacks := make([]stack, 9)
    stackIndex := 0
    
    for i := len(lines)-1; i >= 0; i-- {
        line := lines[i]
        for j := 1; j < len(line); j += 4 {
            if line[j] != ' ' {
                stacks[stackIndex].push(rune(line[j]))
            }
            stackIndex += 1
        }
        stackIndex = 0
    }

    moves := strings.Split(parts[1], "\n")
    
    for _, move := range moves {
        var amount, from, to int
        fmt.Sscanf(move, "move %d from %d to %d", &amount, &from, &to)  

        for i := amount; i > 0; i-- {
            ele := stacks[from-1].pop()
            stacks[to-1].push(ele)
        }
    }

    var result string
    for _, s := range stacks {
       result += string(s.top())
    }

    return result
}

func Part2(filename string) int {
    

    return -1
}
