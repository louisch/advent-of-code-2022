package util

import (
    "bufio"
    "fmt"
    "os"
)

func ReadFile(day int, part int) string {
    filename := fmt.Sprintf("data/%v/%v", day, part)
    data, err := os.ReadFile(filename)
    Check(err)
    return string(data)
}

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
