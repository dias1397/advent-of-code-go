package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
    cycle, register int
}

func main() {
    fmt.Print("--- Day X: TEMPLATE ---\n\n")

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
    registerValues := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}

    cpu := cpu{0, 1}
    for _, line := range strings.Split(input, "\n") {
        instruction := strings.Split(line, " ") 

        cpu.cycle++

        _, ok := registerValues[cpu.cycle]
        if ok {
            registerValues[cpu.cycle] = cpu.cycle * cpu.register 
        }
        
        if instruction[0] == "addx" {
            cpu.cycle++

            value, err := strconv.Atoi(instruction[1])
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error while converting: %v", instruction[1]) 
            }

            _, ok := registerValues[cpu.cycle]
            if ok {
                registerValues[cpu.cycle] = cpu.cycle * cpu.register 
            }

            cpu.register += value 
        }
    }

    result := 0
    for _, value := range registerValues {
        result += value
    }

    return result 
}

func Part2(input string) int {

    return -1 
}
