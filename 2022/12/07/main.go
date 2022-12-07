/*
 */
package main

import (
	"fmt"
	"io"
	"os"
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

	acc, _ := root.Fprint(w)
	var sum int
	for _, n := range acc {
		sum += n
	}
	fmt.Fprintf(w, "Answer: %d\n", sum)
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
		return fmt.Sprintf("%s (dir)", n.name)
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}

func (n *Node) Fprint(w io.Writer) (acc []int, size int) {
	return n.fprint(w, 0)
}

func (n *Node) fprint(w io.Writer, indent int) (acc []int, size int) {
	fmt.Fprintf(w, "%*s %s\n", indent, "-", n)
	switch n.fsType {
	case FSFile:
		return nil, n.size
	case FSDirectory:
		for _, child := range n.children {
			cacc, csize := child.fprint(w, indent+2)
			size += csize
			acc = append(acc, cacc...)
			//size += csize
		}
		var include bool
		if size <= 100000 {
			acc = append(acc, size)
			include = true
		}
		fmt.Fprintf(w, "%*ssize=%d (%t)\n", indent+1, "", size, include)
		fmt.Fprintf(w, "%*sacc=%#v\n", indent+1, "", acc)
		return
	default:
		panic(fmt.Errorf("unrecognized type: %d", n.fsType))
	}
}
