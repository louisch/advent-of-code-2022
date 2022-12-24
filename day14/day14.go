package day14

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type RockMap = map[util.Coords]bool

func readMap(day int) (RockMap, int) {
    rockMap := make(RockMap)
    largestY := 0
    visitLine := func (line string) {
        lineSplit := strings.Split(line, " -> ")

        var lastWall *util.Coords = nil

        for _, coordsSpec := range lineSplit {
            xStr, yStr, ok := strings.Cut(coordsSpec, ",")
            if !ok {
                log.Fatal(fmt.Sprintf("Line '%v' split = '%v', could not find comma!", line, coordsSpec))
                return
            }
            coords := util.Coords {
                X: util.ParseIntSimple(xStr),
                Y: util.ParseIntSimple(yStr),
            }
            if coords.Y > largestY {
                largestY = coords.Y
            }
            if lastWall != nil && lastWall.X == coords.X {
                startY, endY := lastWall.Y, coords.Y
                if coords.Y < lastWall.Y {
                    startY, endY = coords.Y, lastWall.Y
                }
                for Y := startY; Y <= endY; Y++ {
                    rockMap[coords] = true
                }
            } else if lastWall != nil && lastWall.Y == coords.Y {
                startX, endX := lastWall.X, coords.X
                if coords.X < lastWall.X {
                    startX, endX = coords.X, lastWall.X
                }
                for x := startX; x <= endX; x++ {
                    rockMap[util.Coords { X: x, Y: coords.Y }] = true
                }
            } else if lastWall != nil {
                log.Fatal(fmt.Sprintf("Last wall %v is diagonal from %v on line '%v'!", lastWall.Fmt(), coords.Fmt(), line))
                return
            }
            lastWall = &coords
        }
    }
    util.ScanFileByLine(day, visitLine)
    return rockMap, largestY
}

func neXtPosForSand(sandPos util.Coords, rockMap RockMap) *util.Coords {
    coordsBelow := util.Coords { X: sandPos.X, Y: sandPos.Y + 1 }
    if _, hasRock := rockMap[coordsBelow]; !hasRock {
        return &coordsBelow
    }
    coordsLeftBelow := util.Coords { X: sandPos.X - 1, Y: sandPos.Y + 1 }
    if _, hasRock := rockMap[coordsLeftBelow]; !hasRock {
        return &coordsLeftBelow
    }
    coordsRightBelow := util.Coords { X: sandPos.X + 1, Y: sandPos.Y + 1 }
    if _, hasRock := rockMap[coordsRightBelow]; !hasRock {
        return &coordsRightBelow
    }
    return nil
}

func Part1(day int, part int) {
    sandSpawn := util.Coords { X: 500, Y: 0 }
    rockMap, bottomY := readMap(day)

    sandPos := sandSpawn
    sandUnits := 0
    for {
        if sandPos.Y > bottomY {
            break
        }

        neXtSandPos := neXtPosForSand(sandPos, rockMap)
        if neXtSandPos == nil {
            if sandPos.X == sandSpawn.X && sandPos.Y == sandSpawn.Y {
                log.Fatal("Sand is going to block spawn!")
                return
            }
            sandUnits++
            rockMap[sandPos] = true
            sandPos = sandSpawn
            continue
        }
        sandPos = *neXtSandPos
    }

    fmt.Printf("Sand units fallen: %v\n", sandUnits)
}

func Part2(day int, part int) {
    sandSpawn := util.Coords { X: 500, Y: 0 }
    rockMap, bottomY := readMap(day)
    floorY := bottomY + 2

    sandPos := sandSpawn
    sandUnits := 0
    for {
        if _, spawnBlocked := rockMap[sandSpawn]; spawnBlocked {
            break
        }
        neXtSandPos := neXtPosForSand(sandPos, rockMap)
        if neXtSandPos == nil || neXtSandPos.Y >= floorY {
            sandUnits++
            rockMap[sandPos] = true
            sandPos = sandSpawn
            continue
        }
        sandPos = *neXtSandPos
    }

    fmt.Printf("Sand units fallen: %v\n", sandUnits)
}
