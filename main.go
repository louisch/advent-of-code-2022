package main

import (
    "fmt"

    "github.com/louisch/advent-of-code-2022/day1"
)

var challenges = map[string]func(){
    "1:1": day1.Part1,
    "1:2": day1.Part2,
}

func main() {
    fmt.Println("Choose day:")
    var day string
    fmt.Scanln(&day)
    fmt.Println("Choose part:")
    var part string
    fmt.Scanln(&part)

    key := fmt.Sprintf("%v:%v", day, part)
    challenge := challenges[key]
    if challenge == nil {
        fmt.Println("Unknown day %v or part %v chosen!", day, part)
    }
    challenge()
}
