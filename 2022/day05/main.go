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
    fmt.Print("--- Day 5: Supply Stacks ---\n\n")

    if len(os.Args) > 1 {
        input, err := os.ReadFile(os.Args[1]) 
        if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to read file named %s", os.Args[1])
            os.Exit(-1)
        }

        fmt.Fprintf(os.Stdout, "Part1 solution: %s\n", Part1(string(input)))
        fmt.Fprintf(os.Stdout, "Part2 solution: %s\n", Part2(string(input)))
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

func Part1(input string) string {
    parts := strings.Split(input, "\n\n")

    lines := strings.Split(parts[0], "\n")
    lines = lines[:len(lines)-1]

    numberOfStacks := (1 + len(lines[0])) / 4

    stacks := make([]stack, numberOfStacks)
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

func Part2(input string) string {
    parts := strings.Split(input, "\n\n")

    lines := strings.Split(parts[0], "\n")
    lines = lines[:len(lines)-1]

    numberOfStacks := (1 + len(lines[0])) / 4

    stacks := make([]stack, numberOfStacks)
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

        var poppedContainers []rune

        for i := amount; i > 0; i-- {
            poppedContainers = append(poppedContainers, stacks[from-1].pop()) 
        }

        for j := amount; j > 0; j-- {
            stacks[to-1].push(poppedContainers[j-1])
        }
    } 

    var result string
    for _, s := range stacks {
       result += string(s.top())
    }

    return result 
}
