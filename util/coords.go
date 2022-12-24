package util

import (
    "fmt"
)

type Coords struct {
    X int
    Y int
}

func (coords *Coords) Fmt() string {
    return fmt.Sprintf("(%v,%v)", coords.X, coords.Y)
}
