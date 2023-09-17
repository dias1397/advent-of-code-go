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

func Part1(input string) int {
    input = strings.TrimSuffix(input, "\n")
    registerValues := map[int]int{20: 0, 60: 0, 100: 0, 140: 0, 180: 0, 220: 0}

    cpu := cpu{0, 1}
    for _, line := range strings.Split(input, "\n") {
        instruction := strings.Split(line, " ") 

        cpu.cycle++
        updateRegisterValues(registerValues, cpu)
        
        if instruction[0] == "addx" {
            cpu.cycle++
            updateRegisterValues(registerValues, cpu)

            value, err := strconv.Atoi(instruction[1])
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error while converting: %v", instruction[1]) 
                return -1
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

func Part2(input string) string {
    input = strings.TrimSuffix(input, "\n")
    var crtPixels [240]string

    cpu := cpu{0, 1}
    for _, line := range strings.Split(input, "\n") {
        instruction := strings.Split(line, " ") 

        drawPixels(&crtPixels, cpu)
        cpu.cycle++
        
        if instruction[0] == "addx" {
            drawPixels(&crtPixels, cpu)
            cpu.cycle++

            value, err := strconv.Atoi(instruction[1])
            if err != nil {
                fmt.Fprintf(os.Stderr, "Error while converting: %v", instruction[1]) 
                return "-1"
            }

            cpu.register += value 
        }
    }

    result := ""
    for i := 0; i < 6; i++ {
        for j := 0; j < 40; j++ {
            result += crtPixels[i*40+j] 
        }
        result += "\n"
    }

    return result 
}

func updateRegisterValues(registerValues map[int]int, cpu cpu) {
    _, ok := registerValues[cpu.cycle]
    if ok {
        registerValues[cpu.cycle] = cpu.cycle * cpu.register 
    }
}

func drawPixels(pixels *[240]string, cpu cpu) {
    pixelPosition := cpu.cycle % 40
    if pixelPosition <= cpu.register+1 && pixelPosition >= cpu.register-1 {
        pixels[cpu.cycle] = "#" 
    } else {
        pixels[cpu.cycle] = "."
    }
}

func main() {
    fmt.Print("--- Day 10: TEMPLATE ---\n\n")

    if len(os.Args) > 1 {
        input, err := os.ReadFile(os.Args[1]) 
        if err != nil {
            fmt.Fprintf(os.Stderr, "Unable to read file named %s", os.Args[1])
            os.Exit(-1)
        }

        fmt.Fprintf(os.Stdout, "Part1 solution: %d\n", Part1(string(input)))
        fmt.Fprintf(os.Stdout, "Part2 solution: %s\n", Part2(string(input)))
    } else {
        fmt.Fprintf(os.Stderr, "Input file not provided")
    }
}
