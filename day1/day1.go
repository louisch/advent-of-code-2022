package day1

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
)

func scanElves(filename string) []int {
    f, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    var elves = make([]int, 0)
    currentSum := 0
    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            elves = append(elves, currentSum)
            currentSum = 0
            continue
        }

        calories64, err := strconv.ParseInt(line, 0, 32)
        if err != nil {
            log.Fatal(err)
        }
        calories := int(calories64)
        currentSum += calories
    }

    sort.Ints(elves)
    return elves
}

func Part1() {
    elves := scanElves("data/1/1")
    highestCalories := elves[0]
    fmt.Printf("highest calories: %v\n", highestCalories)
}

func Part2() {
    elves := scanElves("data/1/2")
    sum := 0
    toSum := 3
    for i, calories := range elves[len(elves) - toSum:] {
        fmt.Printf("%vth elf has %v\n", i, calories)
        sum += calories
    }
    fmt.Printf("sum of highest %v: %v\n", toSum, sum)
}

func main() {
}
