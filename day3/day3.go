package day3

import (
    "fmt"
    "log"

    "github.com/louisch/advent-of-code-2022/util"
)

var lowercaseStartingPriority = 1
var uppercaseStartingPriority = 27

func getPriority(luggage rune) int {
    intCode := int(luggage)
    if intCode >= int('A') && intCode <= int('Z') {
        return intCode - int('A') + uppercaseStartingPriority
    }
    if intCode >= int('a') && intCode <= int('z') {
        return intCode - int('a') + lowercaseStartingPriority
    }
    log.Fatal(fmt.Sprintf("Given luggage '%v' was not within a known range", luggage))
    return -1
}

func stringToPriorities(str string) map[rune]int {
    priorities := make(map[rune]int)
    for _, c := range str {
        priorities[c] = getPriority(c)
    }
    return priorities
}

func Part1(day int, part int) {
    sum := 0
    visitFile := func(line string) {
        if len(line) % 2 != 0 {
            log.Fatal(fmt.Sprintf("Line %v has length %v which is not even!", line, len(line)))
        }
        compartment1Priorities := stringToPriorities(line[:len(line) / 2])
        var commonPriority int = -1
        for _, c := range line[len(line) / 2:] {
            if priority, exists := compartment1Priorities[c]; exists {
                commonPriority = priority
                break
            }
        }
        if commonPriority < 0 {
            log.Fatal(fmt.Sprintf("Could not find common priorities from line %v!", line))
        }
        sum += commonPriority
    }
    util.ScanFileByLine(day, part, visitFile)
    fmt.Printf("Sum: %v\n", sum)
}

func Part2(day int, part int) {
    linesInThrees := make([][3]string, 0)
    var threeLines [3]string
    i := 0
    visitFile := func(line string) {
        threeLines[i] = line
        if i == 2 {
            linesInThrees = append(linesInThrees, threeLines)
            threeLines = [3]string { "", "", "" }
            i = -1
        }
        i++
    }
    util.ScanFileByLine(day, part, visitFile)

    if i != 0 {
        log.Fatal("Elves are not multiple of three!")
    }

    sum := 0
    GroupFor:
    for i, group := range linesInThrees {
        rucksack1Priorities := stringToPriorities(group[0])
        commonPriorities := make(map[rune]int)
        for _, c := range group[1] {
            if commonPriority, exists := rucksack1Priorities[c]; exists {
                commonPriorities[c] = commonPriority
            }
        }
        for _, c := range group[2] {
            if commonPriority, exists := commonPriorities[c]; exists {
                sum += commonPriority
                continue GroupFor
            }
        }
        errorMessage := fmt.Sprintf("Could not find a priority common to all for group %v!\n", i)
        for _, line := range group {
            errorMessage += fmt.Sprintf("%v\n", line)
        }
        log.Fatal(errorMessage)
    }

    fmt.Printf("Sum: %v\n", sum)
}
