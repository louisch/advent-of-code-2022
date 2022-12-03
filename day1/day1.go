package day1

import (
    "fmt"
    "sort"
    "strconv"

    "github.com/louisch/advent-of-code-2022/util"
)


func scanElves(day int, part int) []int {
    var elves = make([]int, 0)
    currentSum := 0

    visitElves := func(line string) {
        if line == "" {
            elves = append(elves, currentSum)
            currentSum = 0
            return
        }

        calories64, err := strconv.ParseInt(line, 0, 32)
        util.Check(err)
        calories := int(calories64)
        currentSum += calories
    }
    util.ScanFileByLine(day, part, visitElves)

    sort.Ints(elves)
    return elves
}

func Part1(day int, part int) {
    elves := scanElves(day, part)
    highestCalories := elves[0]
    fmt.Printf("highest calories: %v\n", highestCalories)
}

func Part2(day int, part int) {
    elves := scanElves(day, part)
    sum := 0
    toSum := 3
    for i, calories := range elves[len(elves) - toSum:] {
        fmt.Printf("%vth elf has %v\n", i, calories)
        sum += calories
    }
    fmt.Printf("sum of highest %v: %v\n", toSum, sum)
}
