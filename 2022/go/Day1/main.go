package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func main() {
    fmt.Println("Day1 Solution")

    if os.Args[1:] != nil {
        solution := Part2(os.Args[1])
        fmt.Fprintf(os.Stdout, "Solution: %d", solution)
    } else {
        fmt.Fprintf(os.Stderr, "Filename not provided")
    }
}

func Part1(filename string) int {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File not found")
        return -1
    }

    scan := bufio.NewScanner(file)

    var mostCalories = 0
    var currentCalories = 0

    for scan.Scan() {
        if scan.Text() == "" {
            if currentCalories > mostCalories {
                mostCalories = currentCalories
            }        
            currentCalories = 0
        } else {
            calories, err := strconv.Atoi(scan.Text())
            if err != nil {
                fmt.Fprintf(os.Stderr, "Invalid calories value")
            }
            currentCalories += calories
        }
    }

    return mostCalories 
}

func Part2(filename string) int {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File not found")
        return -1
    }

    scan := bufio.NewScanner(file)

    calories := []int{}
    cal := 0

    for scan.Scan() {
        if scan.Text() == "" {
            calories = append(calories, cal)
            cal = 0
        } else {
            calories, err := strconv.Atoi(scan.Text())
            if err != nil {
                fmt.Fprintf(os.Stderr, "Invalid calories value")
            }
            cal += calories
        }
    }

    sort.Slice(calories, func(i, j int) bool {
        return calories[i] > calories[j]
    })

    solution := 0
    for _, value := range calories[:3] {
        solution += value
    }

    return solution
}
