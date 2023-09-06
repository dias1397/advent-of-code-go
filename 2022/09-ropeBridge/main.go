package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    fmt.Print("--- Day 9: TEMPLATE ---\n\n")

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

    headPosition := [2]int{0, 0}
    tailPosition := [2]int{0, 0}
    ropePositionsVisited := make(map[string]bool)
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
            moveEdgeOfRope(&headPosition, directionOfMovement)    
            adjustRopeTailPosition(&tailPosition, headPosition, directionOfMovement)

            position := fmt.Sprint(tailPosition[0], tailPosition[1])
            ropePositionsVisited[position] = true
        }
    }

    return len(ropePositionsVisited)
}

func Part2(input string) int {

    return -1 
}

func moveEdgeOfRope(ropePosition *[2]int, directionOfMovement string) {
    switch directionOfMovement {
    case "R":
        ropePosition[1] += 1
        break;
    case "L":
        ropePosition[1] -= 1
        break;
    case "U":
        ropePosition[0] += 1
        break;
    case "D":
        ropePosition[0] -= 1
        break;
    }
}

func adjustRopeTailPosition(tailPosition *[2]int, headPosition [2]int, directionOfMovement string) {
    columnDifference := headPosition[0] - tailPosition[0]
    rowDifference := headPosition[1] - tailPosition[1]

    if columnDifference <= 1 && columnDifference >= -1 && 
        rowDifference <= 1 && rowDifference >= -1 {
        return 
    }

    if columnDifference == 0 || rowDifference == 0 {
        moveEdgeOfRope(tailPosition, directionOfMovement) 
        return
    }

    moveEdgeOfRope(tailPosition, directionOfMovement)    
    if columnDifference == 1 {
        moveEdgeOfRope(tailPosition, "U") 
        return
    }
    if columnDifference == -1 {
        moveEdgeOfRope(tailPosition, "D") 
        return
    }
    if rowDifference == 1 {
        moveEdgeOfRope(tailPosition, "R") 
        return
    }
    if rowDifference == -1 {
        moveEdgeOfRope(tailPosition, "L") 
        return
    }
}
