package day9

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Knot [2]int

func createKnot() Knot {
    return Knot { 0, 0 }
}

func getSerializedTailPos(tailPos Knot) string {
    return fmt.Sprintf("%v:%v", tailPos[0], tailPos[1])
}

func knotFollow(knot Knot, knotToFollow Knot) Knot {
    axisDiffs := [2]int { knot[0] - knotToFollow[0], knot[1] - knotToFollow[1] }
    if axisDiffs[0] < -1 ||
       axisDiffs[0] > 1 ||
       axisDiffs[1] < -1 ||
       axisDiffs[1] > 1 {
        for axis := 0; axis < 2; axis++ {
            if axisDiffs[axis] < 0 {
                knot[axis]++
            } else if axisDiffs[axis] > 0 {
                knot[axis]--
            }
        }
    }
    return knot
}

func moveKnot(knot Knot, dir string) Knot {
    switch dir {
    case "U":
        knot[1] += 1
    case "D":
        knot[1] -= 1
    case "R":
        knot[0] += 1
    case "L":
        knot[0] -= 1
    default:
        log.Fatal(fmt.Sprintf("Unknown direction %v!", dir))
    }
    return knot
}

func visitSpec(day int, initialKnots []Knot, onMoveStep func([]Knot) []Knot) []Knot {
    knots := initialKnots

    visitLine := func (line string) {
        dir, stepsStr, found := strings.Cut(line, " ")
        if !found {
            log.Fatal(fmt.Sprintf("Could not split line %v by space!", line))
            return
        }
        steps := util.ParseIntSimple(stepsStr)

        for step := 0; step < steps; step++ {
            for i := 0; i < len(knots); i++ {
                if i == 0 {
                    knots[0] = moveKnot(knots[0], dir)
                    continue
                }
                knots[i] = knotFollow(knots[i], knots[i - 1])
                knots = onMoveStep(knots)
            }
        }
    }
    util.ScanFileByLine(day, visitLine)

    return knots
}

func Part1(day int, part int) {
    knots := []Knot { createKnot(), createKnot() }
    visitedTailPositions := make(map[string]Knot)
    visitedTailPositions[getSerializedTailPos(knots[1])] = knots[1]

    onMoveStep := func (knots []Knot) []Knot {
        visitedTailPositions[getSerializedTailPos(knots[1])] = knots[1]
        return knots
    }
    knots = visitSpec(day, knots, onMoveStep)

    fmt.Printf("Tail visited %v locations\n", len(visitedTailPositions))
}

func Part2(day int, part int) {
    knots := make([]Knot, 10)
    for i := 0; i < 10; i++ {
        knots[i] = createKnot()
    }
    visitedTailPositions := make(map[string]Knot)
    visitedTailPositions[getSerializedTailPos(knots[9])] = knots[9]

    onMoveStep := func (knots []Knot) []Knot {
        visitedTailPositions[getSerializedTailPos(knots[9])] = knots[9]
        return knots
    }
    knots = visitSpec(day, knots, onMoveStep)

    fmt.Printf("Tail visited %v locations\n", len(visitedTailPositions))
}
