package main

import (
    "fmt"
    "strconv"

    "github.com/louisch/advent-of-code-2022/util"
    "github.com/louisch/advent-of-code-2022/day1"
    "github.com/louisch/advent-of-code-2022/day2"
    "github.com/louisch/advent-of-code-2022/day3"
)

var challenges = map[string]func(int, int){
    "1:1": day1.Part1,
    "1:2": day1.Part2,
    "2:1": day2.Part1,
    "2:2": day2.Part2,
    "3:1": day3.Part1,
    "3:2": day3.Part2,
}

func main() {
    fmt.Println("Choose day:")
    var dayAsStr string
    fmt.Scanln(&dayAsStr)
    fmt.Println("Choose part:")
    var partAsStr string
    fmt.Scanln(&partAsStr)

    day64, err := strconv.ParseInt(dayAsStr, 0, 32)
    util.Check(err)
    part64, err := strconv.ParseInt(partAsStr, 0, 32)
    util.Check(err)
    day := int(day64)
    part := int(part64)

    key := fmt.Sprintf("%v:%v", day, part)
    challenge := challenges[key]
    if challenge == nil {
        fmt.Printf("Unknown day %v or part %v chosen!\n", day, part)
    }
    challenge(day, part)
}
