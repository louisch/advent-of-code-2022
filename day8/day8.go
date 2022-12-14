package day8

import (
    "fmt"

    "github.com/louisch/advent-of-code-2022/util"
)

type Tree struct {
    height int
    leftVisible bool
    rightVisible bool
    topVisible bool
    bottomVisible bool
    scenicScore int
}

func fmtTree(tree Tree) string {
    return fmt.Sprintf("Height %v Visibility [%v,%v,%v,%v]", tree.height, tree.leftVisible, tree.topVisible, tree.rightVisible, tree.bottomVisible)
}

func isTreeVisible(tree *Tree) bool {
    return tree.leftVisible || tree.rightVisible || tree.topVisible || tree.bottomVisible
}

func readTrees(day int) [][]Tree {
    trees := make([][]Tree, 0)

    visitLine := func (line string) {
        lineOfTrees := make([]Tree, 0)
        for _, c := range line {
            tree := Tree {
                height: util.ParseIntSimple(string(c)),
                leftVisible: true,
                rightVisible: true,
                topVisible: true,
                bottomVisible: true,
                scenicScore: 0,
            }
            lineOfTrees = append(lineOfTrees, tree)
        }
        trees = append(trees, lineOfTrees)
    }
    util.ScanFileByLine(day, visitLine)

    return trees
}

func visitTrees(trees [][]Tree, visitFunc func (tree *Tree, i int, j int)) [][]Tree {
    for i := 0; i < len(trees); i++ {
        for j := 0; j < len(trees[i]); j++ {
            tree := &trees[i][j]
            visitFunc(tree, i, j)
        }
    }
    return trees
}

func markBlockingTrees(trees [][]Tree) [][]Tree {
    markBlockingTreesOne := func (tree *Tree, i int, j int) {
        for iOther := 0; iOther < len(trees); iOther++ {
            if iOther == i {
                continue
            }

            otherTree := &trees[iOther][j]

            if iOther < i && otherTree.height <= tree.height {
                otherTree.bottomVisible = false
            }
            if iOther > i && otherTree.height <= tree.height {
                otherTree.topVisible = false
            }
        }

        for jOther := 0; jOther < len(trees); jOther++ {
            if jOther == j {
                continue
            }

            otherTree := &trees[i][jOther]

            if jOther < j && otherTree.height <= tree.height {
                otherTree.rightVisible = false
            }
            if jOther > j && otherTree.height <= tree.height {
                otherTree.leftVisible = false
            }
        }
    }

    return visitTrees(trees, markBlockingTreesOne)
}

func Part1(day int, part int) {
    trees := readTrees(day)
    trees = markBlockingTrees(trees)

    numOfVisibleTrees := 0
    findVisibleTree := func(tree *Tree, i int, j int) {
        if isTreeVisible(tree) {
            numOfVisibleTrees++
        }
    }
    visitTrees(trees, findVisibleTree)

    fmt.Printf("Number of visible trees: %v\n", numOfVisibleTrees)
}

func markScenicScores(trees [][]Tree) [][]Tree {
    markScenicScoreOne := func (tree *Tree, i int, j int) {
        topViewableTrees := 0
        bottomViewableTrees := 0
        leftViewableTrees := 0
        rightViewableTrees := 0

        for iAbove := i - 1; iAbove >= 0; iAbove-- {
            topViewableTrees++
            if trees[iAbove][j].height >= tree.height {
                break
            }
        }

        for iBelow := i + 1; iBelow < len(trees); iBelow++ {
            bottomViewableTrees++
            if trees[iBelow][j].height >= tree.height {
                break
            }
        }

        for jLeft := j - 1; jLeft >= 0; jLeft-- {
            leftViewableTrees++
            if trees[i][jLeft].height >= tree.height {
                break
            }
        }

        for jRight := j + 1; jRight < len(trees[i]); jRight++ {
            rightViewableTrees++
            if trees[i][jRight].height >= tree.height {
                break
            }
        }

        tree.scenicScore = topViewableTrees * bottomViewableTrees * leftViewableTrees * rightViewableTrees
    }

    return visitTrees(trees, markScenicScoreOne)
}

func Part2(day int, part int) {
    trees := readTrees(day)
    trees = markScenicScores(trees)

    maxScenicScore := 0
    findMaxScenicScore := func (tree *Tree, i int, j int) {
        if tree.scenicScore > maxScenicScore {
            maxScenicScore = tree.scenicScore
        }
    }
    visitTrees(trees, findMaxScenicScore)

    fmt.Printf("Max Scenic Score: %v\n", maxScenicScore)
}
