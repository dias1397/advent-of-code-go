package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) > 1 {
        solution := Part2(os.Args[1])
        fmt.Fprintf(os.Stdout, "%d", solution)
    }
}

func Part1(filename string) int {
    file, err := os.Open(filename)
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open file")
    }

    defer file.Close()

    scan := bufio.NewScanner(file)

    result := 0
    scores := map[string]int{"X": 1, "Y": 2, "Z": 3, 
                             "A": 1, "B": 2, "C": 3}

    for scan.Scan() {
        if scan.Text() == "" {
            continue
        }

        round := strings.Fields(scan.Text())
        result += scores[round[1]]

        switch (scores[round[1]] - scores[round[0]] + 3) % 3 {
            case 0:
                result += 3
            case 1:
                result += 6
        }
    }

    return result
}

func Part2(filename string) int {
    file, err := os.Open(filename)
    
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to open file")
    }

    defer file.Close()

    scan := bufio.NewScanner(file)
    
    result := 0
    scores := map[string]int{"X": 0, "Y": 3, "Z": 6, 
                             "A": 1, "B": 2, "C": 3}

    for scan.Scan() {
        if scan.Text() == "" {
            continue
        }

        round := strings.Fields(scan.Text())
        result += scores[round[1]]

        var temp int

        switch round[1] {
            case "X":
                temp = (scores[round[0]] - 1) % 3 
            case "Y":
                temp = scores[round[0]]
            case "Z":
                temp = (scores[round[0]] + 1) % 3
        }
        
        if temp == 0 {
            temp = 3
        }

        result += temp
    }

    return result
}
