package main

import (
    "testing"

    "github.com/dias1397/advent-of-code-go/2022/utils"
)

func TestPart1(t *testing.T) {
    var tests = []struct{
        name string
        input string
        want int
    } {
        {"Part1 sample", testutil.ReadInput(t, "./testdata/sample.txt"), 13140},
        {"Part1 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), 14220},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ans := Part1(tt.input)

            if ans != tt.want{
                t.Errorf("Part1 = %d; want %d", ans, tt.want)
            }
        })
    }
}

func TestPart2(t *testing.T) {
    var tests = []struct{
        name string
        input string
        want string
    } {
        {"Part2 sample", testutil.ReadInput(t, "./testdata/sample.txt"), testutil.ReadInput(t, "./testdata/sampleSolution.txt")},
        {"Part2 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), testutil.ReadInput(t, "./testdata/exerciseSolution.txt")},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ans := Part2(tt.input)

            if ans != tt.want{
                t.Errorf("Part2 = %s; want %s", ans, tt.want)
            }
        })
    }
}
