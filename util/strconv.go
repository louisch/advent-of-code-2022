package util

import(
    "strconv"
)

func ParseIntSimple(intSpec string) int {
    result64, err := strconv.ParseInt(intSpec, 10, 32)
    Check(err)
    return int(result64)
}
