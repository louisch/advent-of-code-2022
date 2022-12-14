package day7

import (
    "fmt"
)

func Part1(day int, part int) {
    env := buildEnv(day)

    sum := 0
    sumIfLeq100K := func (node *Inode, _ int) {
        if node.size <= 100000 && node.isDir {
            sum += node.size
        }
    }
    visitEnv(env, sumIfLeq100K)
    fmt.Printf("Total size of directories <= 100K in size: %v\n", sum)
}

func Part2(day int, part int) {
    env := buildEnv(day)

    totalSpace    := 70000000
    requiredSpace := 30000000
    unusedSpace := totalSpace - env.root.size
    needToDelete := requiredSpace - unusedSpace
    if needToDelete < 0 {
        fmt.Printf("Root has size %v, leaving us with %v which is enough space for the %v required for the update!\n", env.root.size, unusedSpace, requiredSpace)
        return
    }
    smallestDirLargeEnough := env.root
    findSmallestDirLargeEnough := func (node *Inode, _ int) {
        if node.isDir && node.size >= needToDelete && node.size < smallestDirLargeEnough.size {
            smallestDirLargeEnough = node
        }
    }
    visitEnv(env, findSmallestDirLargeEnough)
    fmt.Printf("Need to delete %v, smallest dir that is large enough: %v size: %v\n", needToDelete, smallestDirLargeEnough.name, smallestDirLargeEnough.size)
}
