package day7

import (
    "fmt"
    "log"
    "strings"

    "github.com/louisch/advent-of-code-2022/util"
)

type Env struct {
    root *Inode
    currentDir *Inode
    lsMode bool
}

func processCdCmd(cdDir string, env Env) Env {
    switch cdDir {
    case "/":
        env.currentDir = env.root
    case "..":
        env.currentDir = env.currentDir.parent
        if env.currentDir == nil {
            log.Fatal(fmt.Sprintf("%v has no parent!", env.currentDir.name))
        }
    default:
        env.currentDir = navigateToDir(env.currentDir, cdDir)
        if env.currentDir == nil {
            log.Fatal(fmt.Sprintf("Could not find dir %v to navigate to from %v!", cdDir, env.currentDir.name))
        }
    }
    return env
}

func processLsCmd(env Env) Env {
    env.lsMode = true
    return env
}

func processCmd(args []string, env Env) Env {
    env.lsMode = false
    switch args[0] {
    case "cd":
        return processCdCmd(args[1], env)
    case "ls":
        return processLsCmd(env)
    default:
        log.Fatal(fmt.Sprintf("Unknown command %v!", args[0]))
    }
    return env
}

func processLsSpec(line string, env Env) Env {
    nodedesc, name, _ := strings.Cut(line, " ")
    for _, child := range env.currentDir.children {
        if child.name == name {
            return env
        }
    }

    newNode := makeNodeFromLsSpec(nodedesc, name)
    env.currentDir.children = append(env.currentDir.children, &newNode)
    newNode.parent = env.currentDir
    if !newNode.isDir {
        addSizeToAncestors(&newNode)
    }
    return env
}

func addSizeToAncestors(file *Inode) {
    if file.isDir {
        log.Fatal("Only call this function on file nodes!")
        return
    }

    currentNode := file
    for currentNode.parent != nil {
        currentNode = currentNode.parent
        if !currentNode.isDir {
            log.Fatal(fmt.Sprintf("One of this file's ancestors is a file?? File: %v Ancestor: %v", file.name, currentNode.name))
        }
        currentNode.size += file.size
    }
}

func processTreeSpecLine(line string, env Env) Env {
    lineSplit := strings.Split(line, " ")

    if lineSplit[0] == "$" {
        return processCmd(lineSplit[1:], env)
    }
    if env.lsMode {
        return processLsSpec(line, env)
    }

    log.Fatal(fmt.Sprintf("Unknown format for line %v!", line))
    return env
}

func buildEnv(day int) Env {
    root := makeDir("/")
    env := Env {
        root: &root,
        currentDir: &root,
        lsMode: false,
    }
    visitFile := func (line string) {
        env = processTreeSpecLine(line, env)
    }
    util.ScanFileByLine(day, visitFile)
    return env
}

type VisitNode struct {
    node *Inode
    depth int
}

func visitEnv(env Env, nodeFunc func(*Inode, int)) Env {
    toVisit := []VisitNode {
        {
            node: env.root,
            depth: 0,
        },
    }
    for len(toVisit) > 0 {
        current := toVisit[0]
        nodeFunc(current.node, current.depth)
        toVisit = toVisit[1:]
        for _, child := range current.node.children {
            toVisit = append([]VisitNode {
                {
                    node: child,
                    depth: current.depth + 1,
                },
            }, toVisit...)
        }
    }
    return env
}

func debugEnv(env Env) {
    debugNode := func (node *Inode, depth int) {
        indentation := ""
        for i := 0; i < depth; i++ {
            indentation += "  "
        }
        nodeDesc := fmt.Sprintf("dir, total size=%v", node.size)
        if !node.isDir {
            nodeDesc = fmt.Sprintf("file, size=%v", node.size)
        }
        fmt.Printf("%v- %v (%v)\n", indentation, node.name, nodeDesc)
    }
    visitEnv(env, debugNode)
}
