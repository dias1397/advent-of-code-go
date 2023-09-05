package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
    headPosition := [2]int{0, 0}
    tailPosition := [2]int{0, 0}
    ropePositionsVisited := make([]string, 0)
    movements := strings.Split(input, "\n")

    for _, movement := range movements {
        movementInstructions := strings.Split(movement, " ")

        directionOfMovement := movementInstructions[0]
        numberOfMovements, err := strconv.Atoi(movementInstructions[1])
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error while parsing number of movements: %v", err) 
            os.Exit(-1)
        }

        for i := 0; i < numberOfMovements; i++ {
            moveHead(headPosition, directionOfMovement)    
            moveTail(tailPosition, headPosition)

            addPositionToVisitedSet(tailPosition)
        }
    }

    return -1 
}

func Part2(input string) int {

    return -1 
}
