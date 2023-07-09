package main

import (
    "testing"

    "github.com/dias1397/advent-of-code-go/2022/utils"
)

func TestPart1(t *testing.T) {
    var tests = []struct{
        name string
        input string
        want string 
    } {
        {"Part1 sample", testutil.ReadInput(t, "./testdata/sample.txt"), "CMZ"},
        {"Part1 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), "BWNCQRMDB"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            ans := Part1(tt.input)

            if ans != tt.want{
                t.Errorf("Part1 = %s; want %s", ans, tt.want)
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
        {"Part2 sample", testutil.ReadInput(t, "./testdata/sample.txt"), "MCD"},
        {"Part2 exercise", testutil.ReadInput(t, "./testdata/exercise.txt"), "NHWZCBNBF"},
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
