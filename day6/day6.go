package day6

import (
    "fmt"

    "github.com/louisch/advent-of-code-2022/util"
)

func lastNUnique(str string, n int) bool {
    for i := 0; i < n; i++ {
        for j := 1; j < n; j++ {
            if i == j {
                continue
            }
            if str[i] == str[j] {
                return false
            }
        }
    }
    return true
}

func findUniqueCharsNLength(dataStream string, n int) (int, string) {
    lastNChars := dataStream[:n - 1]
    for i := 4; i < len(dataStream); i++ {
        lastNChars += string(dataStream[i])

        if lastNUnique(lastNChars, n) {
            return i, lastNChars
        }

        lastNChars = lastNChars[1:]
    }
    return -1, ""
}

func Part1(day int, part int) {
    dataStream := util.ReadFile(day)
    i, unique4 := findUniqueCharsNLength(dataStream, 4)
    if i >= 0 {
        fmt.Printf("4 unique characters found at %vth position: %v\n", i + 1, unique4)
    } else {
        fmt.Println("Could not find 4 unique characters!")
    }
}

func Part2(day int, part int) {
    dataStream := util.ReadFile(day)
    i, unique14 := findUniqueCharsNLength(dataStream, 14)
    if i >= 0 {
        fmt.Printf("14 unique characters found at %vth position: %v\n", i + 1, unique14)
    } else {
        fmt.Println("Could not find 14 unique characters!")
    }
}
