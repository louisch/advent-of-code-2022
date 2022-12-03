package main

import (
    "fmt"
    "strconv"

    "github.com/louisch/advent-of-code-2022/util"
    "github.com/louisch/advent-of-code-2022/day1"
)

var challenges = map[string]func(int, int){
    "1:1": day1.Part1,
    "1:2": day1.Part2,
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
