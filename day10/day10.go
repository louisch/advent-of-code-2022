package day10

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Instruction struct {
    kind string
    arg int
    cyclesRem int
}

func fmtInstruction(instruction Instruction) string {
    if instruction.kind == "noop" {
        return "noop"
    }
    return fmt.Sprintf("%v %v", instruction.kind, instruction.arg)
}

type Env struct {
    cycle int
    x int
}

func parseLine(line string) *Instruction {
    lineSplit := strings.Split(line, " ")
    switch lineSplit[0] {
    case "noop":
        return &Instruction {
            kind: "noop",
            cyclesRem: 1,
        }
    case "addx":
        return &Instruction {
            kind: "addx",
            arg: util.ParseIntSimple(lineSplit[1]),
            cyclesRem: 2,
        }
    }
    log.Fatal(fmt.Sprintf("Unknown instruction %v!", line))
    return nil
}

func executeInstructions(day int, duringFunc func(Env), condFunc func(Env) bool, boundFunc func(Env) bool) {
    filename := fmt.Sprintf("data/%v", day)
    f, err := os.Open(filename)
    util.Check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    env := Env {
        cycle: 0,
        x: 1,
    }

    var executingInstruction *Instruction = nil

    for {
        if boundFunc(env) {
            break
        }

        //fmt.Printf("< C:%3v X:%3v", env.cycle + 1, env.x)

        if executingInstruction == nil {
            found := scanner.Scan()
            if found {
                executingInstruction = parseLine(scanner.Text())
                //fmt.Printf(" begin %v", executingInstruction.kind)
            } else {
                //fmt.Printf(" end read  ")
            }
        } else {
            //fmt.Printf(" exing %v", executingInstruction.kind)
        }

        if condFunc(env) {
            duringFunc(env)
        }

        if executingInstruction != nil {
            executingInstruction.cyclesRem--
            if executingInstruction.cyclesRem == 0 {
                //fmt.Printf(" execute")
                switch executingInstruction.kind {
                case "addx":
                    env.x += executingInstruction.arg
                }
                executingInstruction = nil
            } else {
                //fmt.Printf(" rem %3v", executingInstruction.cyclesRem)
            }
        } else {
            //fmt.Printf(" no inst")
        }

        //fmt.Printf(" > X:%3v\n", env.x)

        env.cycle++
    }
}

func Part1(day int, part int) {
    limit := 220
    checkAtBeginning := 20
    checkAtEvery := 40
    sum := 0
    duringFunc := func (env Env) {
        signalStrength := (env.cycle + 1) * env.x
        fmt.Printf("Cycle: %v, X: %v, Signal Strength: %v\n", env.cycle + 1, env.x, signalStrength)
        sum += signalStrength
    }
    condFunc := func (env Env) bool {
        return (env.cycle + 1 + checkAtBeginning) % checkAtEvery == 0
    }
    boundFunc := func (env Env) bool {
        return env.cycle > limit
    }
    executeInstructions(day, duringFunc, condFunc, boundFunc)
    fmt.Printf("Total Signal Strength: %v\n", sum)
}

func Part2(day int, part int) {
    line := ""
    currentLine := 0
    screenWidth := 40
    screenHeight := 6
    crtPos := 0
    duringFunc := func (env Env) {
        spritePos := env.x
        if spritePos - 1 <= crtPos && crtPos <= spritePos + 1 {
            line += "#"
        } else {
            line += "."
        }

        crtPos++

        if crtPos >= screenWidth {
            fmt.Println(line)
            line = ""
            currentLine++
            crtPos = 0
        }
    }
    condFunc := func (env Env) bool {
        return true
    }
    boundFunc := func (env Env) bool {
        return env.cycle >= screenWidth * screenHeight
    }
    executeInstructions(day, duringFunc, condFunc, boundFunc)
}
