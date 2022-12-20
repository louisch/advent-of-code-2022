package day11

import(
    "fmt"
    "log"
    "regexp"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Item = int

type Monkey struct {
    items []Item
    operation [3]string
    divisor int
    test [2]int
    inspections int
}

type MonkeyBuilder struct {
    startingItems []Item
    operation *[3]string
    divisor *int
    test *[2]int
}

func readMonkeys(day int) ([]Monkey, int) {
    monkeys := make([]Monkey, 0)
    builder := MonkeyBuilder {
        startingItems: make([]Item, 0),
        operation: nil,
        test: nil,
    }
    commonMultiple := 1

    buildMonkey := func () {
        if builder.operation == nil {
            log.Fatal("Builder has no operation!")
            return
        }
        if builder.test == nil {
            log.Fatal("Builder has no test!")
            return
        }
        commonMultiple *= *builder.divisor
        monkeys = append(monkeys, Monkey {
            items: builder.startingItems,
            operation: *builder.operation,
            divisor: *builder.divisor,
            test: *builder.test,
            inspections: 0,
        })

        builder = MonkeyBuilder {
            startingItems: make([]Item, 0),
            operation: nil,
            divisor: nil,
            test: nil,
        }
    }

    visitLine := func (line string) {
        // Match blank line
        if line == "" {
            return
        }

        // Match monkey index
        matched, err := regexp.MatchString(`Monkey \d+:`, line)
        util.Check(err)
        if matched {
            return
        }

        kind, rest, found := strings.Cut(strings.TrimSpace(line), ": ")
        if !found {
            log.Fatal(fmt.Sprintf("Non-blank line %v had no colon!", line))
            return
        }

        switch kind {
        case "Starting items":
            items := strings.Split(rest, ", ")
            for _, itemStr := range items {
                item := util.ParseIntSimple(itemStr)
                builder.startingItems = append(builder.startingItems, item)
            }
        case "Operation":
            regex, err := regexp.Compile(`new = (\S+) (\S) (\S+)`)
            util.Check(err)
            submatches := regex.FindStringSubmatch(rest)
            builder.operation = &[3]string { submatches[1], submatches[2], submatches[3] }
        case "Test":
            regex, err := regexp.Compile(`divisible by (\d+)`)
            util.Check(err)
            submatches := regex.FindStringSubmatch(rest)
            divisor := util.ParseIntSimple(submatches[1])
            builder.divisor = &divisor
        case "If true":
            regex, err := regexp.Compile(`throw to monkey (\d+)`)
            util.Check(err)
            submatches := regex.FindStringSubmatch(rest)
            targetMonkey := util.ParseIntSimple(submatches[1])
            builder.test = &[2]int { targetMonkey, 0 }
        case "If false":
            regex, err := regexp.Compile(`throw to monkey (\d+)`)
            util.Check(err)
            submatches := regex.FindStringSubmatch(rest)
            targetMonkey := util.ParseIntSimple(submatches[1])
            builder.test[1] = targetMonkey

            buildMonkey() // If false should be the last part of a monkey spec
        default:
            log.Fatal(fmt.Sprintf("Unrecognized kind of line %v", line))
        }
    }
    util.ScanFileByLine(day, visitLine)

    return monkeys, commonMultiple
}

func processMonkeysOneRound(monkeys []Monkey, decreaseWorryFactor int, commonMultiple int) []Monkey {
    for monkeyIndex := 0; monkeyIndex < len(monkeys); monkeyIndex++ {
        monkey := monkeys[monkeyIndex]
        for itemIndex := 0; itemIndex < len(monkey.items); itemIndex++ {
            worryLevel := monkey.items[itemIndex]
            monkey.inspections++

            parseArg := func (arg string) int {
                if arg == "old" {
                    return worryLevel
                }
                return util.ParseIntSimple(arg)
            }

            arg1 := parseArg(monkey.operation[0])
            arg2 := parseArg(monkey.operation[2])
            switch monkey.operation[1] {
            case "+":
                worryLevel = arg1 + arg2
            case "*":
                worryLevel = arg1 * arg2
            }

            worryLevel /= decreaseWorryFactor
            worryLevel %= commonMultiple

            targetMonkey := monkey.test[1]
            if worryLevel % monkey.divisor == 0 {
                targetMonkey = monkey.test[0]
            }
            if targetMonkey == monkeyIndex {
                log.Fatal("Cannot support throwing to self!")
                return []Monkey {}
            }
            monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, worryLevel)
        }
        monkey.items = []Item {}
        monkeys[monkeyIndex] = monkey
    }
    return monkeys
}

func printMonkeys(monkeys []Monkey) {
    for i, monkey := range monkeys {
        fmt.Printf("Monkey %v inspected %v, items:", i, monkey.inspections)
        for j, item := range monkey.items {
            if j != 0 {
                fmt.Print(",")
            }
            fmt.Printf(" %v", item)
        }
        fmt.Println()
    }
}

func calculateMonkeyBusiness(day int, decreaseWorryFactor int, rounds int) int {
    monkeys, commonMultiple := readMonkeys(day)
    for i := 0; i < rounds; i++ {
        monkeys = processMonkeysOneRound(monkeys, decreaseWorryFactor, commonMultiple)
    }

    highestInspection := 0
    secondHighestInspection := 0
    for _, monkey := range monkeys {
        if monkey.inspections > highestInspection {
            secondHighestInspection = highestInspection
            highestInspection = monkey.inspections
        } else if monkey.inspections > secondHighestInspection {
            secondHighestInspection = monkey.inspections
        }
    }

    return highestInspection * secondHighestInspection
}

func Part1(day int, part int) {
    monkeyBusiness := calculateMonkeyBusiness(11, 3, 20)
    fmt.Printf("Monkey business: %v\n", monkeyBusiness)
}

func Part2(day int, part int) {
    monkeyBusiness := calculateMonkeyBusiness(11, 1, 10000)
    fmt.Printf("Monkey business: %v\n", monkeyBusiness)
}
