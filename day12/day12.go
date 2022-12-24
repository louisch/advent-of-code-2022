package day12

import (
    "fmt"
    "log"

    "github.com/louisch/advent-of-code-2022/util"
)

type Coords struct {
    x int
    y int
}

type Elevation struct {
    c rune
    visited bool
}

type ElevationMap struct {
    start *Coords
    end *Coords
    tiles [][]Elevation
    width int
    height int
}

func readElevationMap(day int) *ElevationMap {
    var start *Coords = nil
    var end *Coords = nil
    tiles := make([][]Elevation, 0)
    y := 0
    width := -1
    height := 0
    visitLine := func (line string) {
        if width < 0 {
            width = len(line)
        }

        if len(line) != width {
            log.Fatal("One line is a different width to the rest!")
        }

        row := make([]Elevation, width)
        for x, c := range line {
            if c == 'S' {
                start = &Coords {
                    x,
                    y,
                }
            }
            if c == 'E' {
                end = &Coords {
                    x,
                    y,
                }
            }
            row[x] = Elevation {
                c: c,
                visited: false,
            }
        }
        tiles = append(tiles, row)
        y++
        height++
    }
    util.ScanFileByLine(day, visitLine)

    return &ElevationMap {
        start,
        end,
        tiles,
        width,
        height,
    }
}

func validNeighbors(coords Coords, width int, height int) []Coords {
    neighbors := []Coords {
        {
            coords.x - 1,
            coords.y,
        },
        {
            coords.x + 1,
            coords.y,
        },
        {
            coords.x,
            coords.y - 1,
        },
        {
            coords.x,
            coords.y + 1,
        },
    }
    validNeighbors := make([]Coords, 0)
    for _, neighbor := range neighbors {
        if neighbor.x >= 0 && neighbor.x < width &&
           neighbor.y >= 0 && neighbor.y < height {
            validNeighbors = append(validNeighbors, neighbor)
        }
    }
    return validNeighbors
}

func findPath(elevationMap *ElevationMap, start Coords) (int, bool) {
    boundary := map[Coords]Elevation {
        start: elevationMap.tiles[elevationMap.start.y][elevationMap.start.x],
    }
    lengthOfPath := -1
    pathFound := false
    for !pathFound && len(boundary) > 0 {
        lengthOfPath++
        newBoundary := make(map[Coords]Elevation, 0)
        for coords, current := range boundary {
            elevationMap.tiles[coords.y][coords.x].visited = true
            currentElevation := current.c
            if currentElevation == 'S' {
                currentElevation = 'a'
            }
            if currentElevation == 'E' {
                pathFound = true
                break
            }

            visitableCoords := validNeighbors(coords, elevationMap.width, elevationMap.height)
            for _, neighborCoords := range visitableCoords {
                neighbor := elevationMap.tiles[neighborCoords.y][neighborCoords.x]
                neighborElevation := neighbor.c
                if neighborElevation == 'E' {
                    neighborElevation = 'z'
                }
                _, exists := newBoundary[neighborCoords]
                if !neighbor.visited && neighborElevation - currentElevation <= 1 && !exists {
                    newBoundary[neighborCoords] = neighbor
                }
            }
        }
        boundary = newBoundary
    }
    return lengthOfPath, pathFound
}

func resetMap(elevationMap *ElevationMap) {
    for y := 0; y < len(elevationMap.tiles); y++ {
        for x := 0; x < len(elevationMap.tiles[y]); x++ {
            elevationMap.tiles[y][x].visited = false
        }
    }
}

func Part1(day int, part int) {
    elevationMap := readElevationMap(day)
    if elevationMap.start == nil {
        log.Fatal("No start found on elevation map!")
        return
    }
    start := *elevationMap.start
    lengthOfPath, pathFound := findPath(elevationMap, start)

    if pathFound {
        fmt.Printf("Length of path: %v\n", lengthOfPath)
    } else {
        fmt.Println("Path to top not found!")
    }
}

func Part2(day int, part int) {
    elevationMap := readElevationMap(day)
    if elevationMap.start == nil {
        log.Fatal("No start found on elevation map!")
        return
    }
    startLocations := map[Coords]Elevation {
        *elevationMap.start: elevationMap.tiles[elevationMap.start.y][elevationMap.start.x],
    }
    for y := 0; y < len(elevationMap.tiles); y++ {
        for x := 0; x < len(elevationMap.tiles[y]); x++ {
            coords := Coords { x, y }
            _, existing := startLocations[coords]
            elevation := elevationMap.tiles[y][x]
            if !existing && elevation.c == 'a' {
                startLocations[coords] = elevation
            }
        }
    }

    smallestPath := -1
    for coords := range startLocations {
        lengthOfPath, pathFound := findPath(elevationMap, coords)
        if pathFound &&
            (smallestPath < 0 || lengthOfPath < smallestPath) {
            smallestPath = lengthOfPath
        }
        resetMap(elevationMap)
    }

    if smallestPath > 0 {
        fmt.Printf("Length of path: %v\n", smallestPath)
    } else {
        fmt.Println("Path to top not found!")
    }
}
