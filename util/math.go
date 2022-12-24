package util

func Abs(n int) int {
    if n < 0 {
        n = -n
    }
    return n
}

type Range struct {
    Lower int
    Upper int
}

func (aRange *Range) Len() int {
    return aRange.Upper - aRange.Lower
}

func (aRange *Range) Contains(n int) bool {
    return aRange.Lower <= n && n <= aRange.Upper
}

func (aRange *Range) Overlaps(otherRange *Range) bool {
    return aRange.Contains(otherRange.Lower) || aRange.Contains(otherRange.Upper) ||
        otherRange.Contains(aRange.Lower) || otherRange.Contains(aRange.Upper)
}

func (aRange *Range) Merge(otherRange *Range) *Range {
    if otherRange.Lower < aRange.Lower {
        aRange.Lower = otherRange.Lower
    }
    if otherRange.Upper > aRange.Upper {
        aRange.Upper = otherRange.Upper
    }
    return aRange
}

func (aRange *Range) Subtract(otherRange *Range) *Range {
    if aRange.Contains(otherRange.Lower) {
        aRange.Upper = otherRange.Lower - 1
    }
    if aRange.Contains(otherRange.Upper) {
        aRange.Lower = otherRange.Upper + 1
    }
    if otherRange.Contains(aRange.Lower) || otherRange.Contains(aRange.Upper) {
        return nil
    }
    return aRange
}

func (aRange *Range) IsValid() bool {
    return aRange != nil && aRange.Lower <= aRange.Upper
}
