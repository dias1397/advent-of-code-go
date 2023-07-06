package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
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

    for _, caloriesByElf := range strings.Split(input, "\n\n") {
        sumOfElfCalories := 0

        for _, calories := range strings.Fields(caloriesByElf) {
            temp, err := strconv.Atoi(calories)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Unable to convert %s to int", calories)
                os.Exit(-1)
            }
            sumOfElfCalories += temp
        }

        if sumOfElfCalories > result {
            result = sumOfElfCalories
        }
    }

    return result 
}

func Part2(input string) int {
    result := 0

    sumOfCaloriesForAllElves := []int{}

    for _, caloriesByElf := range strings.Split(input, "\n\n") {
        sumOfElfCalories := 0

        for _, calories := range strings.Fields(caloriesByElf) {
            temp, err := strconv.Atoi(calories)
            if err != nil {
                fmt.Fprintf(os.Stderr, "Unable to convert %s to int", calories)
                os.Exit(-1)
            }
            sumOfElfCalories += temp
        }

        sumOfCaloriesForAllElves = append(sumOfCaloriesForAllElves, sumOfElfCalories) 
    }

    sort.Slice(sumOfCaloriesForAllElves, func(i, j int) bool {
        return sumOfCaloriesForAllElves[i] > sumOfCaloriesForAllElves[j]
    })

    for _, value := range sumOfCaloriesForAllElves[:3] {
        result += value
    }

    return result 
}
