package day7

import (
    "fmt"
    "log"

    "github.com/louisch/advent-of-code-2022/util"
)

type Inode struct {
    name string
    size int
    isDir bool
    children []*Inode
    parent *Inode
}

func fmtInode(node Inode) string {
    return fmt.Sprintf("name: %v, size: %v, isDir: %v, numOfChildren: %v", node.name, node.size, node.isDir, len(node.children))
}

func navigateToDir(currentNode *Inode, dirToNavigate string) *Inode {
    if !currentNode.isDir {
        msg := fmt.Sprintf("Attempt to navigate from a non-directory %v to %v!", fmtInode(*currentNode), dirToNavigate)
        log.Fatal(msg)
    }

    for _, child := range currentNode.children {
        if child.name == dirToNavigate {
            return child
        }
    }

    return nil
}

func makeDir(name string) Inode {
    return Inode {
        name: name,
        size: 0,
        isDir: true,
        children: make([]*Inode, 0),
        parent: nil,
    }
}

func makeFile(name string, size int) Inode {
    return Inode {
        name: name,
        size: size,
        isDir: false,
        children: make([]*Inode, 0),
        parent: nil,
    }
}

func makeNodeFromLsSpec(nodedesc string, name string) Inode {
    if nodedesc == "dir" {
        return makeDir(name)
    }
    size := util.ParseIntSimple(nodedesc)
    return makeFile(name, size)
}
