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
        {"Part1 sample", testutil.ReadInput(t, "./testdata/sample.txt"), -1},
        {"Part1 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), -1},
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
        want int
    } {
        {"Part2 sample", testutil.ReadInput(t, "./testdata/sample.txt"), -1},
        {"Part2 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ans := Part2(tt.input)

            if ans != tt.want{
                t.Errorf("Part2 = %d; want %d", ans, tt.want)
            }
        })
    }
}
