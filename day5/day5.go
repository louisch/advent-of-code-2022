package day5

import (
    "fmt"
    "regexp"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type CrateInstruction = struct {
    amount int
    from int
    to int
}

type CrateSpec = struct {
    stacks []string
    instructions []CrateInstruction
}

func parseSpec(day int) CrateSpec {
    stacks := make([]string, 0)
    instructions := make([]CrateInstruction, 0)

    stackLines := make([]string, 0)
    blankLineReached := false
    visitFile := func(line string) {
        // Once all stacks are read, parse the read stackLines into stacks
        if line == "" {
            // Convert read stackLines into stacks
            stackNumbers := strings.Split(stackLines[len(stackLines) - 1], " ")
            lastStackNumber := stackNumbers[len(stackNumbers) - 1]
            numberOfStacks := util.ParseIntSimple(lastStackNumber)
            stacks = make([]string, numberOfStacks)
            for i := 0; i < numberOfStacks; i++ {
                stacks[i] = ""
            }
            // Read the rest of the stackLines to get the actual stacks
            for i := len(stackLines) - 2; i >= 0; i-- {
                stackLine := stackLines[i]
                for j := 0; j < numberOfStacks; j++ {
                    indexOfCrate := 1 + j * 4
                    if indexOfCrate >= len(stackLine) {
                        break
                    }
                    crate := string(stackLine[indexOfCrate])
                    if crate == " " {
                        continue
                    }
                    stacks[j] += crate
                }
            }

            blankLineReached = true
            return
        }

        // Reading stacks at this point (saving into stackLines for parsing later)
        if !blankLineReached {
            stackLines = append(stackLines, line)
            return
        }

        // Parsing instructions from this point onwards
        regex, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
        util.Check(err)
        submatches := regex.FindStringSubmatch(line)
        instructions = append(instructions, CrateInstruction {
            amount: util.ParseIntSimple(submatches[1]),
            from: util.ParseIntSimple(submatches[2]) - 1,
            to: util.ParseIntSimple(submatches[3]) - 1,
        })
    }
    util.ScanFileByLine(day, visitFile)

    return CrateSpec {
        stacks,
        instructions,
    }
}

func printMessage(stacks []string) {
    message := ""
    for _, stack := range stacks {
        message += string(stack[len(stack) - 1])
    }
    fmt.Printf("Message: %v\n", message)
}

func Part1(day int, part int) {
    spec := parseSpec(day)
    stacks := spec.stacks
    for _, instruction := range spec.instructions {
        amount := instruction.amount
        from := instruction.from
        to := instruction.to
        cratesToMove := stacks[from][len(stacks[from]) - amount:]
        stacks[from] = stacks[from][:len(stacks[from]) - amount]
        for i := len(cratesToMove) - 1; i >= 0; i-- {
            stacks[to] += string(cratesToMove[i])
        }
    }

    printMessage(stacks)
}

func Part2(day int, part int) {
    spec := parseSpec(day)
    stacks := spec.stacks
    for _, instruction := range spec.instructions {
        amount := instruction.amount
        from := instruction.from
        to := instruction.to
        cratesToMove := stacks[from][len(stacks[from]) - amount:]
        stacks[from] = stacks[from][:len(stacks[from]) - amount]
        for _, crate := range cratesToMove {
            stacks[to] += string(crate)
        }
    }

    printMessage(stacks)
}
