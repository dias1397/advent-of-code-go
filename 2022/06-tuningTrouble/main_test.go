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
        {"Part1 sample 1", testutil.ReadInput(t, "./testdata/sample1.txt"), 7},
        {"Part1 sample 2", testutil.ReadInput(t, "./testdata/sample2.txt"), 5},
        {"Part1 sample 3", testutil.ReadInput(t, "./testdata/sample3.txt"), 6},
        {"Part1 sample 4", testutil.ReadInput(t, "./testdata/sample4.txt"), 10},
        {"Part1 sample 5", testutil.ReadInput(t, "./testdata/sample5.txt"), 11},
        {"Part1 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), 1802},
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
        {"Part2 sample 1", testutil.ReadInput(t, "./testdata/sample1.txt"), 19},
        {"Part2 sample 2", testutil.ReadInput(t, "./testdata/sample2.txt"), 23},
        {"Part2 sample 3", testutil.ReadInput(t, "./testdata/sample3.txt"), 23},
        {"Part2 sample 4", testutil.ReadInput(t, "./testdata/sample4.txt"), 29},
        {"Part2 sample 5", testutil.ReadInput(t, "./testdata/sample5.txt"), 26},
        {"Part2 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), 3551},
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
