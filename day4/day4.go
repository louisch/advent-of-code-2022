package day4

import (
    "fmt"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

func getRanges(line string) (int, int, int, int) {
    range1, range2, _ := strings.Cut(line, ",")
    range1LowerStr, range1UpperStr, _ := strings.Cut(range1, "-")
    range2LowerStr, range2UpperStr, _ := strings.Cut(range2, "-")
    range1Lower := util.ParseIntSimple(range1LowerStr)
    range1Upper := util.ParseIntSimple(range1UpperStr)
    range2Lower := util.ParseIntSimple(range2LowerStr)
    range2Upper := util.ParseIntSimple(range2UpperStr)
    return range1Lower, range1Upper, range2Lower, range2Upper
}

func doRangesIntersectCompletely(lower1 int, upper1 int, lower2 int, upper2 int) bool {
    return lower1 >= lower2 && upper1 <= upper2 ||
           lower2 >= lower1 && upper2 <= upper1
}

func doRangesNotIntersect(lower1 int, upper1 int, lower2 int, upper2 int) bool {
    return upper1 < lower2 || lower1 > upper2
}

func Part1(day int, part int) {
    rangesThatIntersectCompletely := 0
    visitFile := func(line string) {
        lower1, upper1, lower2, upper2 := getRanges(line)
        if (doRangesIntersectCompletely(lower1, upper1, lower2, upper2)) {
            rangesThatIntersectCompletely++
        }
    }
    util.ScanFileByLine(day, visitFile)
    fmt.Printf("Ranges intersecting completely: %v\n", rangesThatIntersectCompletely)
}

func Part2(day int, part int) {
    rangesThatIntersectPartially := 0
    visitFile := func(line string) {
        lower1, upper1, lower2, upper2 := getRanges(line)
        if (!doRangesNotIntersect(lower1, upper1, lower2, upper2)) {
            rangesThatIntersectPartially++
        }
    }
    util.ScanFileByLine(day, visitFile)
    fmt.Printf("Ranges intersecting partially: %v\n", rangesThatIntersectPartially)
}
