package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
    x, y int
}

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

    headPosition := point{0, 0}
    tailPosition := point{0, 0}
    tailPositionVisited := make(map[point]bool)

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
            moveRopeHead(&headPosition, directionOfMovement)
            adjustRopeTail(&tailPosition, headPosition)

            tailPositionVisited[tailPosition] = true
        }
    }

    return len(tailPositionVisited)
}

func Part2(input string) int {
    input = strings.TrimSuffix(input, "\n")

    ropeKnotsPositions := make([]point, 10)
    tailPositionVisited := make(map[point]bool)

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
            moveRopeHead(&ropeKnotsPositions[0], directionOfMovement)

            for j := 1; j < 10; j++ {
                adjustRopeTail(&ropeKnotsPositions[j], ropeKnotsPositions[j-1])
            }

            tailPositionVisited[ropeKnotsPositions[9]] = true
        }
    }


    return len(tailPositionVisited) 
}

func moveRopeHead(ropePosition *point, directionOfMovement string) {
    switch directionOfMovement {
    case "U":
        (*ropePosition).x += 1
    case "D":
        (*ropePosition).x -= 1
    case "R":
        (*ropePosition).y += 1
    case "L":
        (*ropePosition).y -= 1
    }
}

func adjustRopeTail(tailPosition *point, headPosition point) {
    rowDifference := float64(headPosition.x - (*tailPosition).x)
    columnDifference := float64(headPosition.y - (*tailPosition).y)

    if math.Abs(rowDifference) == 2 ||
        (math.Abs(rowDifference) == 1 && math.Abs(columnDifference) == 2) {
        (*tailPosition).x += int(rowDifference) / int(math.Abs(rowDifference))
    }

    if math.Abs(columnDifference) == 2 ||
        (math.Abs(columnDifference) == 1 && math.Abs(rowDifference) == 2) {
        (*tailPosition).y += int(columnDifference) / int(math.Abs(columnDifference))
    }
}
