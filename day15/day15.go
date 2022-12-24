package day15

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Sensor struct {
    coords util.Coords
    closestBeacon *Beacon
}

type Beacon struct {
    coords util.Coords
}

func pairSorted(n1, n2 int) (int, int) {
    if n1 > n2 {
        n1, n2 = n2, n1
    }
    return n1, n2
}

func readSensors(day int) []Sensor {
    sensors := make([]Sensor, 0)
    visitLine := func (line string) {
        pattern := `Sensor at x=(-?\d*), y=(-?\d*): closest beacon is at x=(-?\d*), y=(-?\d*)`
        regex := util.RegexpCompileSimple(pattern)
        submatches := regex.FindStringSubmatch(strings.TrimSpace(line))
        if submatches == nil {
            log.Fatal(fmt.Sprintf("Line: \n%v\ndoes not match pattern:\n%v", line, pattern))
        }

        sensorX := util.ParseIntSimple(submatches[1])
        sensorY := util.ParseIntSimple(submatches[2])
        closestBeaconX := util.ParseIntSimple(submatches[3])
        closestBeaconY := util.ParseIntSimple(submatches[4])

        beacon := Beacon {
            coords: util.Coords { X: closestBeaconX, Y: closestBeaconY },
        }
        sensorCoords := util.Coords { X: sensorX, Y: sensorY }
        sensors = append(sensors, Sensor {
            coords: sensorCoords,
            closestBeacon: &beacon,
        })
    }
    util.ScanFileByLine(day, visitLine)
    return sensors
}

func occupiedRanges(sensors []Sensor, y int, countBeaconsAndSensors bool) []*util.Range {
    ranges := make([]*util.Range, 0)

    for _, sensor := range sensors {
        closestBeacon := sensor.closestBeacon
        distX := util.Abs(sensor.coords.X - closestBeacon.coords.X)
        distY := util.Abs(sensor.coords.Y - closestBeacon.coords.Y)
        rangeRadius := distX + distY
        distYToSensor := util.Abs(sensor.coords.Y - y)
        rangeYLow  := sensor.coords.Y - rangeRadius
        rangeYHigh := sensor.coords.Y + rangeRadius

        if y < rangeYLow || rangeYHigh < y || distYToSensor > rangeRadius {
            continue
        }

        rangeXLow  := sensor.coords.X - rangeRadius + distYToSensor
        rangeXHigh := sensor.coords.X + rangeRadius - distYToSensor

        proposedRanges := make([]*util.Range, 0)
        if !countBeaconsAndSensors && distYToSensor == 0 {
            proposedRanges = []*util.Range {
                { Lower: rangeXLow, Upper: sensor.coords.X - 1 },
                { Lower: sensor.coords.X + 1, Upper: rangeXHigh },
            }
        } else {
            proposedRanges = []*util.Range {
                { Lower: rangeXLow, Upper: rangeXHigh },
            }
        }

        for _, proposed := range proposedRanges {
            makeNewRange := true
            for _, aRange := range ranges {
                if aRange.Overlaps(proposed) {
                    aRange.Merge(proposed)
                    makeNewRange = false
                    break
                }
            }

            if makeNewRange {
                ranges = append(ranges, proposed)
            }
        }

        newRanges := make([]*util.Range, 0)
        if len(ranges) > 1 {
            for len(ranges) > 0 {
                aRange := ranges[0]
                ranges = ranges[1:]
                for _, otherRange := range ranges {
                    if aRange.Overlaps(otherRange) {
                        otherRange.Merge(aRange)
                        aRange = nil
                        break
                    }
                }
                if aRange != nil {
                    newRanges = append(newRanges, aRange)
                }
            }
            ranges = newRanges
        }
    }

    return ranges
}

func Part1(day int, part int) {
    sensors := readSensors(day)

    testRow := 2000000
    occupiedRanges := occupiedRanges(sensors, testRow, false)

    if len(occupiedRanges) == 0 {
        fmt.Println("Could not find any ranges??")
    } else {
        occupiedSpaces := 0
        for _, aRange := range occupiedRanges {
            occupiedSpaces += aRange.Len()
        }
        fmt.Printf("Number of positions without beacon: %v\n", occupiedSpaces)
    }

}

func Part2(day int, part int) {
    sensors := readSensors(day)

    found := false
    searchLower := 0
    searchUpper := 4000000
    for y := searchLower; y <= searchUpper; y++ {
        unoccupiedRange := &util.Range {
            Lower: searchLower,
            Upper: searchUpper,
        }

        occupiedRanges := occupiedRanges(sensors, y, true)
        for _, occupied := range occupiedRanges {
            unoccupiedRange = unoccupiedRange.Subtract(occupied)
            if !unoccupiedRange.IsValid() {
                break
            }
        }

        if unoccupiedRange.IsValid() {
            found = true
            if unoccupiedRange.Len() > 1 {
                log.Fatal("Found more than one place for distress signal!")
                return
            }
            tuningFrequency := unoccupiedRange.Lower * 4000000 + y
            fmt.Printf("Distress Signal at %v,%v, Tuning Frequency: %v\n", unoccupiedRange.Lower, y, tuningFrequency)
            break
        }
    }

    if !found {
        fmt.Println("Could not find distress signal")
    }
}

