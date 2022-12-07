/*
 */
package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	framework "github.com/ppg/advent-of-code/2022/12/framework"
)

func main() {
	framework.Register(parser, solution0)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func solution0(w io.Writer, runner *framework.Runner[string]) {
	root := &Node{name: "/", children: make(map[string]*Node), fsType: FSDirectory}
	cur := root
	for line := range runner.Lines() {
		switch line[0:1] {
		case "$": // command
			cols := strings.Split(line, " ")
			cmd := cols[1]
			switch cmd {
			case "cd":
				name := cols[2]
				switch name {
				case "/":
					fmt.Fprintf(w, "changing to root directory\n")
					cur = root
				case "..":
					fmt.Fprintf(w, "changing to parent directory\n")
					cur = cur.parent
				default:
					fmt.Fprintf(w, "changing to %s directory\n", name)
					cur = cur.children[name]
				}
			case "ls":
				// do nothing here?
				//panic("unimplemented")
			default:
				panic(fmt.Errorf("unrecognized command: %s (%s)", cmd, line))
			}

		default: // ls output
			cols := strings.Split(line, " ")
			name := cols[1]
			switch cols[0] {
			case "dir":
				fmt.Fprintf(w, "found directory %s\n", name)
				cur.children[name] = &Node{parent: cur, name: name, fsType: FSDirectory, children: make(map[string]*Node)}
			default: // file
				size, _ := strconv.Atoi(cols[0])
				fmt.Fprintf(w, "found file %s: %d\n", name, size)
				cur.children[name] = &Node{parent: cur, name: name, fsType: FSFile, size: size}
			}
		}
	}

	// Print the tree, accumulating the dirs and the total FS size
	// TODO(ppg): rename from Fprintf
	dirs, size := root.Fprint(w)

	// Calculate required space
	unused := 70000000 - size
	fmt.Fprintf(w, "    unused: %d\n", unused)
	required := 30000000 - unused
	fmt.Fprintf(w, "  required: %d\n", required)

	var sum int
	var chosen *Node
	fmt.Fprintf(w, "candidates\n")

	// Go through directories from largest to smallest; find
	// - (part1) the sum of all directories < 100000
	// - (part2) the smallest dir that is greater than required size
	// TODO(ppg): use heap to avoid sorting
	sort.Stable(BySize(dirs))
	for _, dir := range dirs {
		fmt.Fprintf(w, "  %s\n", dir)
		if dir.size <= 100000 {
			sum += dir.size
		}
		if dir.size > required {
			chosen = dir
		}
	}
	fmt.Fprintf(w, "sum (part 1): %d\n", sum)
	fmt.Fprintf(w, "chosen (part 2):\n  %s\n", chosen)
}

type FSType int

const (
	FSFile FSType = iota
	FSDirectory
)

func (t FSType) String() string {
	switch t {
	case FSFile:
		return "file"
	case FSDirectory:
		return "dir"
	default:
		panic(fmt.Errorf("unrecognized type: %d", t))
	}
}

type BySize []*Node

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return a[i].size > a[j].size }

type Node struct {
	parent   *Node
	name     string
	children map[string]*Node
	fsType   FSType
	size     int
}

func (n *Node) String() string {
	switch n.fsType {
	case FSFile:
		return fmt.Sprintf("%s (file, size=%d)", n.name, n.size)
	case FSDirectory:
		if n.size > 0 {
			return fmt.Sprintf("%s (dir, size=%d)", n.name, n.size)
		} else {
			return fmt.Sprintf("%s (dir)", n.name)
		}
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}

func (n *Node) Fprint(w io.Writer) (acc []*Node, size int) {
	return n.fprint(w, 0)
}

func (n *Node) fprint(w io.Writer, indent int) (acc []*Node, size int) {
	fmt.Fprintf(w, "%*s %s\n", indent, "-", n)
	switch n.fsType {
	case FSFile:
		return nil, n.size
	case FSDirectory:
		for _, child := range n.children {
			cacc, csize := child.fprint(w, indent+2)
			size += csize
			acc = append(acc, cacc...)
		}
		// store the size we computed for this node
		n.size = size
		acc = append(acc, n)
		fmt.Fprintf(w, "%*ssize=%d\n", indent+1, "", size)
		return
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}
