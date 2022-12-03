package util

import (
    "bufio"
    "fmt"
    "os"
)

func ScanFileByLine(day int, part int, visit func(string)) {
    filename := fmt.Sprintf("data/%v/%v", day, part)
    f, err := os.Open(filename)
    Check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        visit(line)
    }
}
