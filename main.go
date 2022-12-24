package main

import (
    "fmt"
    "strconv"

    "github.com/louisch/advent-of-code-2022/util"
    "github.com/louisch/advent-of-code-2022/day1"
    "github.com/louisch/advent-of-code-2022/day2"
    "github.com/louisch/advent-of-code-2022/day3"
    "github.com/louisch/advent-of-code-2022/day4"
    "github.com/louisch/advent-of-code-2022/day5"
    "github.com/louisch/advent-of-code-2022/day6"
    "github.com/louisch/advent-of-code-2022/day7"
    "github.com/louisch/advent-of-code-2022/day8"
    "github.com/louisch/advent-of-code-2022/day9"
    "github.com/louisch/advent-of-code-2022/day10"
    "github.com/louisch/advent-of-code-2022/day11"
    "github.com/louisch/advent-of-code-2022/day12"
    "github.com/louisch/advent-of-code-2022/day13"
    "github.com/louisch/advent-of-code-2022/day14"
    "github.com/louisch/advent-of-code-2022/day15"
    "github.com/louisch/advent-of-code-2022/day16"
    "github.com/louisch/advent-of-code-2022/day17"
    "github.com/louisch/advent-of-code-2022/day18"
    "github.com/louisch/advent-of-code-2022/day19"
    "github.com/louisch/advent-of-code-2022/day20"
    "github.com/louisch/advent-of-code-2022/day21"
    "github.com/louisch/advent-of-code-2022/day22"
    "github.com/louisch/advent-of-code-2022/day23"
    "github.com/louisch/advent-of-code-2022/day24"
    "github.com/louisch/advent-of-code-2022/day25"
)

var challenges = map[string]func(int, int){
    "1:1": day1.Part1,
    "1:2": day1.Part2,
    "2:1": day2.Part1,
    "2:2": day2.Part2,
    "3:1": day3.Part1,
    "3:2": day3.Part2,
    "4:1": day4.Part1,
    "4:2": day4.Part2,
    "5:1": day5.Part1,
    "5:2": day5.Part2,
    "6:1": day6.Part1,
    "6:2": day6.Part2,
    "7:1": day7.Part1,
    "7:2": day7.Part2,
    "8:1": day8.Part1,
    "8:2": day8.Part2,
    "9:1": day9.Part1,
    "9:2": day9.Part2,
    "10:1": day10.Part1,
    "10:2": day10.Part2,
    "11:1": day11.Part1,
    "11:2": day11.Part2,
    "12:1": day12.Part1,
    "12:2": day12.Part2,
    "13:1": day13.Part1,
    "13:2": day13.Part2,
    "14:1": day14.Part1,
    "14:2": day14.Part2,
    "15:1": day15.Part1,
    "15:2": day15.Part2,
    "16:1": day16.Part1,
    "16:2": day16.Part2,
    "17:1": day17.Part1,
    "17:2": day17.Part2,
    "18:1": day18.Part1,
    "18:2": day18.Part2,
    "19:1": day19.Part1,
    "19:2": day19.Part2,
    "20:1": day20.Part1,
    "20:2": day20.Part2,
    "21:1": day21.Part1,
    "21:2": day21.Part2,
    "22:1": day22.Part1,
    "22:2": day22.Part2,
    "23:1": day23.Part1,
    "23:2": day23.Part2,
    "24:1": day24.Part1,
    "24:2": day24.Part2,
    "25:1": day25.Part1,
    "25:2": day25.Part2,
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
