/*
 */
package main

import (
	"container/heap"
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
	framework.Register(parser, solution1)
	framework.Run(os.Stdout)
}

var parser = framework.LineParser

func parseFS(w io.Writer, runner *framework.Runner[string]) *Node {
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
	return root
}

func solution0(w io.Writer, runner *framework.Runner[string]) {
	root := parseFS(w, runner)

	// Print the tree, accumulating the dirs and the total FS size
	dirs, size := root.solution0(w, 0)

	// Calculate required space
	unused := 70000000 - size
	fmt.Fprintf(w, "    unused: %d\n", unused)
	required := 30000000 - unused
	fmt.Fprintf(w, "  required: %d\n", required)

	// Go through directories from largest to smallest; find
	// - (part1) the sum of all directories < 100000
	// - (part2) the smallest dir that is greater than required size
	// TODO(ppg): use heap to avoid sorting
	var sum int
	var chosen *Node
	sort.Stable(framework.Array[*Node](dirs))
	fmt.Fprintf(w, "candidates\n")
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

func solution1(w io.Writer, runner *framework.Runner[string]) {
	root := parseFS(w, runner)

	// Print the tree, accumulating the dirs and the total FS size
	dirs, size := root.solution1(w, 0)

	// Calculate required space
	unused := 70000000 - size
	fmt.Fprintf(w, "    unused: %d\n", unused)
	required := 30000000 - unused
	fmt.Fprintf(w, "  required: %d\n", required)

	// Go through directories (heap has them smallest to largest); find
	// - (part1) the sum of all directories < 100000
	// - (part2) the smallest dir that is greater than required size
	var sum int
	var chosen *Node
	fmt.Fprintf(w, "candidates\n")
	for dirs.Len() > 0 {
		dir := heap.Pop(&dirs).(*Node)
		fmt.Fprintf(w, "  %s\n", dir)
		if dir.size <= 100000 {
			sum += dir.size
		}
		if dir.size > required && chosen == nil {
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

// Less sorts largest size to smallest size.
func (n *Node) Less(other *Node) bool { return n.size < other.size }

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
		}
		return fmt.Sprintf("%s (dir)", n.name)
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}

func (n *Node) solution0(w io.Writer, indent int) (dirs []*Node, size int) {
	fmt.Fprintf(w, "%*s %s\n", indent, "-", n)
	switch n.fsType {
	case FSFile:
		return nil, n.size
	case FSDirectory:
		for _, child := range n.children {
			cdirs, csize := child.solution0(w, indent+2)
			size += csize
			dirs = append(dirs, cdirs...)
		}
		// store the size we computed for this node
		n.size = size
		dirs = append(dirs, n)
		fmt.Fprintf(w, "%*ssize=%d\n", indent+1, "", size)
		return
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}

func (n *Node) solution1(w io.Writer, indent int) (dirs framework.Heap[*Node], size int) {
	fmt.Fprintf(w, "%*s %s\n", indent, "-", n)
	switch n.fsType {
	case FSFile:
		return nil, n.size
	case FSDirectory:
		for _, child := range n.children {
			cdirs, csize := child.solution1(w, indent+2)
			size += csize
			// TODO(ppg): look into a merge function for framework.Heap if it can be
			// done more efficiently
			for cdirs.Len() > 0 {
				heap.Push(&dirs, heap.Pop(&cdirs))
			}
		}
		// store the size we computed for this node
		n.size = size
		heap.Push(&dirs, n)
		fmt.Fprintf(w, "%*ssize=%d\n", indent+1, "", size)
		return
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}
