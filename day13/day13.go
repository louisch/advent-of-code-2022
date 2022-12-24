package day13

import (
    "fmt"
    "log"
    "sort"

    "github.com/louisch/advent-of-code-2022/util"
)

type PacketData struct {
    integer int
    list []*PacketData
}

type Pair struct {
    left PacketData
    right PacketData
}

func isInteger(data *PacketData) bool {
    return data.integer >= 0
}

func newIntegerPD(integer int) *PacketData {
    return &PacketData {
        integer: integer,
        list: []*PacketData {},
    }
}

func newListPD(values ...*PacketData) *PacketData {
    return &PacketData {
        integer: -1,
        list: values,
    }
}

func compare(left *PacketData, right *PacketData) int {
    if isInteger(left) && isInteger(right) {
        if left.integer < right.integer {
            return 1
        } else if left.integer > right.integer {
            return -1
        } else {
            return 0
        }
    }

    if !isInteger(left) && !isInteger(right) {
        i := 0
        for {
            if i >= len(left.list) && i >= len(right.list) {
                return 0
            }
            if i >= len(left.list) {
                return 1
            }
            if i >= len(right.list) {
                return -1
            }

            leftData := left.list[i]
            rightData := right.list[i]
            if leftData == nil || rightData == nil {
                log.Fatal("Cannot have nil inside packet data list!")
                return 0
            }
            comparison := compare(leftData, rightData)
            if comparison != 0 {
                return comparison
            }

            i++
        }
    }

    if isInteger(left) {
        return compare(newListPD(left), right)
    }
    return compare(left, newListPD(right))
}

func parseLine(line string) *PacketData {
    var root *PacketData = nil
    var parents = []*PacketData {}
    var currentNumber string = ""

    for _, c := range line {
        var currentList *PacketData = nil
        if len(parents) > 0 {
            currentList = parents[len(parents) - 1]
        }
        if c >= '0' && c <= '9' {
            currentNumber += string(c)
            continue
        }
        if currentNumber != "" {
            newNumber := newIntegerPD(util.ParseIntSimple(currentNumber))
            currentNumber = ""
            if root == nil {
                root = newNumber
            } else if currentList != nil {
                currentList.list = append(currentList.list, newNumber)
            } else {
                log.Fatal("root is not number but no list to append to??")
                return nil
            }
        }
        if c == ',' {
            continue
        }

        if c == '[' {
            newList := newListPD()
            if root == nil {
                root = newList
            }
            if len(parents) > 0 {
                currentList.list = append(currentList.list, newList)
            }
            parents = append(parents, newList)

            continue
        }

        if c == ']' {
            if len(parents) == 0 {
                log.Fatal("Too many closing brackets!")
                return nil
            }
            currentList = parents[len(parents) - 1]
            parents = parents[:len(parents) - 1]
            continue
        }

        log.Fatal(fmt.Sprintf("Unknown character %v in line %v!", string(c), line))
    }

    if len(parents) > 0 {
        log.Fatal("Missing bracket for a [")
        return nil
    }

    return root
}

func readPD(day int) PacketDataArr {
    packetData := []*PacketData {}
    visitLine := func (line string) {
        if line == "" {
            return
        }
        parsedLine := parseLine(line)
        if parsedLine == nil {
            log.Fatal(fmt.Sprintf("Could not parse line! %v", line))
            return
        }
        packetData = append(packetData, parsedLine)
    }
    util.ScanFileByLine(day, visitLine)
    return packetData
}

func Part1(day int, part int) {
    pd := readPD(day)

    if len(pd) % 2 != 0 {
        log.Fatal("There must be an even number of packet data!")
        return
    }

    sum := 0
    for i := 0; i * 2 < len(pd); i++ {
        left := pd[i * 2]
        right := pd[i * 2 + 1]
        comparison := compare(left, right)
        if comparison == 0 {
            fmt.Println("Warning: pair compared as same")
        }
        if comparison == 1 {
            sum += i + 1
        }
    }

    fmt.Printf("Sum: %v\n", sum)
}

type PacketDataArr []*PacketData
func (a PacketDataArr) Len() int           { return len(a) }
func (a PacketDataArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PacketDataArr) Less(i, j int) bool { return compare(a[i], a[j]) == 1 }

func isDivider(pd *PacketData) bool {
    return !isInteger(pd) &&
        len(pd.list) == 1 && !isInteger(pd.list[0]) &&
        len(pd.list[0].list) == 1 && isInteger(pd.list[0].list[0]) &&
        (pd.list[0].list[0].integer == 2 || pd.list[0].list[0].integer == 6)
}

func createDividers() (*PacketData, *PacketData) {
    divider1 := newListPD(newListPD(newIntegerPD(2)))
    divider2 := newListPD(newListPD(newIntegerPD(6)))
    return divider1, divider2
}

func Part2(day int, part int) {
    allPD := readPD(day)
    divider1, divider2 := createDividers()
    allPD = append(allPD, divider1)
    allPD = append(allPD, divider2)
    sort.Sort(allPD)

    decoderKey := 1
    for i, pd := range allPD {
        if isDivider(pd) {
            decoderKey *= i + 1
        }
    }

    fmt.Printf("Decoder Key: %v\n", decoderKey)
}
