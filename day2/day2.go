package day2

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type RPS = int

const (
    Rock RPS = 0
    Paper    = 1
    Scissors = 2
)

var opponentRPSEncoding = map[string]RPS {
    "A": Rock,
    "B": Paper,
    "C": Scissors,
}

var playerRPSEncodingPart1 = map[string]RPS {
    "X": Rock,
    "Y": Paper,
    "Z": Scissors,
}

var playerRPSEncodingPart2 = map[string]int {
    "X": 0,
    "Y": 1,
    "Z": 2,
}

func RPSToString(rps RPS) string {
    switch (rps) {
    case Rock:
        return "Rock"
    case Scissors:
        return "Scissors"
    case Paper:
        return "Paper"
    }
    return "Unknown"
}

type roundGuide struct {
    opponent RPS
    play RPS
    result int
}

func getEncoding(table map[string]int, input string) int {
    result, ok := table[input]
    if !ok {
        log.Fatal(fmt.Sprintf("unknown spec found: %v", input))
    }
    return result
}

func scanRPS(day int, part int) []roundGuide {
    var spec = make([]roundGuide, 0)

    visitFile := func(line string) {
        opponentEncoded, playEncoded, found := strings.Cut(line, " ")
        if !found {
            log.Fatal(fmt.Sprintf("line has no space: %v", line))
        }
        opponent := getEncoding(opponentRPSEncoding, opponentEncoded)
        play := getEncoding(playerRPSEncodingPart1, playEncoded)
        result := getEncoding(playerRPSEncodingPart2, playEncoded)
        spec = append(spec, roundGuide {
            opponent,
            play,
            result,
        })
    }
    util.ScanFileByLine(day, part, visitFile)

    return spec
}

func getOutcome(player RPS, opponent RPS) int {
    diff := int(player) - int(opponent)

    // Rock = 0, Paper = 1, Scissors = 2
    // R
    // 0 -1 -2
    // D  L  W
    // P
    // 1  0 -1
    // W  D  L
    // S
    // 2  1  0
    // L  W  D

    if opponent < 0 || opponent > 2 {
        log.Fatal(fmt.Sprintf("Unknown opponent spec: %v vs %v", player, opponent))
    }

    switch player {
    case Rock:
        if diff == -2 {
            return 1
        }
        return diff
    case Paper:
        return diff
    case Scissors:
        if diff == 2 {
            return -1
        }
        return diff
    }

    log.Fatal(fmt.Sprintf("Unknown player spec: %v vs %v", player, opponent))
    return diff
}

func getScore(guide roundGuide) int {
    player, opponent := guide.play, guide.opponent

    shapeScore := 1
    switch player {
    case Paper:
        shapeScore = 2
    case Scissors:
        shapeScore = 3
    }

    outcome := getOutcome(player, opponent)
    outcomeScore := (outcome + 1) * 3

    return shapeScore + outcomeScore
}

func Part1(day int, part int) {
    spec := scanRPS(day, part)
    score := 0
    for _, guide := range spec {
        score += getScore(guide)
    }
    fmt.Printf("Total Score: %v\n", score)
}

func getGuidedPlay(guide roundGuide) RPS {
    opponent := guide.opponent
    desiredOutcome := guide.result

    // 0 = L, 1 = D, 2 = W
    // R = 0
    // S = 2, R = 0, P = 1
    // P = 1
    // R = 0, P = 1, S = 2
    // S = 2
    // P = 1, S = 2, R = 0

    switch opponent {
    case Rock:
        desiredOutcome -= 1
        if desiredOutcome < 0 {
            return 2
        }
        return desiredOutcome
    case Paper:
        return desiredOutcome
    case Scissors:
        desiredOutcome += 1
        if desiredOutcome > 2 {
            return 0
        }
        return desiredOutcome
    }

    log.Fatal(fmt.Sprintf("Unknown opponent spec: %v", opponent))
    return 0
}

func Part2(day int, part int) {
    spec := scanRPS(day, part)
    score := 0
    for _, guide := range spec {
        play := getGuidedPlay(guide)
        guide = roundGuide {
            opponent: guide.opponent,
            play: play,
            result: guide.result,
        }
        score += getScore(guide)
    }
    fmt.Printf("Total Score: %v\n", score)
}
