package util

import (
    "bufio"
    "fmt"
    "os"
)

func ReadFile(day int) string {
    filename := fmt.Sprintf("data/%v", day)
    data, err := os.ReadFile(filename)
    Check(err)
    return string(data)
}

func ScanFileByLine(day int, visit func(string)) {
    filename := fmt.Sprintf("data/%v", day)
    f, err := os.Open(filename)
    Check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        visit(line)
    }
}
