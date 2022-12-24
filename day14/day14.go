package day14

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Coords struct {
    x int
    y int
}

type RockMap = map[Coords]bool

func readMap(day int) (RockMap, int) {
    rockMap := make(RockMap)
    largestY := 0
    visitLine := func (line string) {
        lineSplit := strings.Split(line, " -> ")

        var lastWall *Coords = nil

        for _, coordsSpec := range lineSplit {
            xStr, yStr, ok := strings.Cut(coordsSpec, ",")
            if !ok {
                log.Fatal(fmt.Sprintf("Line '%v' split = '%v', could not find comma!", line, coordsSpec))
                return
            }
            x := util.ParseIntSimple(xStr)
            y := util.ParseIntSimple(yStr)
            if y > largestY {
                largestY = y
            }
            coords := Coords { x, y }
            if lastWall != nil && lastWall.x == x {
                startY, endY := lastWall.y, y
                if y < lastWall.y {
                    startY, endY = y, lastWall.y
                }
                for y := startY; y <= endY; y++ {
                    rockMap[Coords { x, y }] = true
                }
            } else if lastWall != nil && lastWall.y == y {
                startX, endX := lastWall.x, x
                if x < lastWall.x {
                    startX, endX = x, lastWall.x
                }
                for x := startX; x <= endX; x++ {
                    rockMap[Coords { x, y }] = true
                }
            } else if lastWall != nil {
                log.Fatal(fmt.Sprintf("Last wall (%v,%v) is diagonal from (%v,%v) on line '%v'!", lastWall.x, lastWall.y, x, y, line))
                return
            }
            lastWall = &coords
        }
    }
    util.ScanFileByLine(day, visitLine)
    return rockMap, largestY
}

func nextPosForSand(sandPos Coords, rockMap RockMap) *Coords {
    coordsBelow := Coords { x: sandPos.x, y: sandPos.y + 1 }
    if _, hasRock := rockMap[coordsBelow]; !hasRock {
        return &coordsBelow
    }
    coordsLeftBelow := Coords { x: sandPos.x - 1, y: sandPos.y + 1 }
    if _, hasRock := rockMap[coordsLeftBelow]; !hasRock {
        return &coordsLeftBelow
    }
    coordsRightBelow := Coords { x: sandPos.x + 1, y: sandPos.y + 1 }
    if _, hasRock := rockMap[coordsRightBelow]; !hasRock {
        return &coordsRightBelow
    }
    return nil
}

func Part1(day int, part int) {
    sandSpawn := Coords { x: 500, y: 0 }
    rockMap, bottomY := readMap(day)

    sandPos := sandSpawn
    sandUnits := 0
    for {
        if sandPos.y > bottomY {
            break
        }

        nextSandPos := nextPosForSand(sandPos, rockMap)
        if nextSandPos == nil {
            if sandPos.x == sandSpawn.x && sandPos.y == sandSpawn.y {
                log.Fatal("Sand is going to block spawn!")
                return
            }
            sandUnits++
            rockMap[sandPos] = true
            sandPos = sandSpawn
            continue
        }
        sandPos = *nextSandPos
    }

    fmt.Printf("Sand units fallen: %v\n", sandUnits)
}

func Part2(day int, part int) {
    sandSpawn := Coords { x: 500, y: 0 }
    rockMap, bottomY := readMap(day)
    floorY := bottomY + 2

    sandPos := sandSpawn
    sandUnits := 0
    for {
        if _, spawnBlocked := rockMap[sandSpawn]; spawnBlocked {
            break
        }
        nextSandPos := nextPosForSand(sandPos, rockMap)
        if nextSandPos == nil || nextSandPos.y >= floorY {
            sandUnits++
            rockMap[sandPos] = true
            sandPos = sandSpawn
            continue
        }
        sandPos = *nextSandPos
    }

    fmt.Printf("Sand units fallen: %v\n", sandUnits)
}
