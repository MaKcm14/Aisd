package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type NodeSide int16

const (
	None     NodeSide = 0
	LeftSide NodeSide = iota + 1
	RightSide
	Root
)

type (
	Node struct {
		Key    int64
		Value  string
		Left   *Node
		Right  *Node
		Parent *Node
	}

	Tree struct {
		Root *Node
	}
)

func NewTree() Tree {
	return Tree{}
}

func New(root *Node) Tree {
	return Tree{
		Root: root,
	}
}

func (tree *Tree) find(key int64) *Node {
	for startVert := tree.Root; startVert != nil; {
		if startVert.Key > key {
			startVert = startVert.Left
		} else if startVert.Key < key {
			startVert = startVert.Right
		} else {
			return startVert
		}
	}
	return nil
}

func (tree *Tree) defineParentSide(node *Node) (NodeSide, error) {
	if node != nil {
		if node == tree.Root {
			return Root, nil
		} else if node.Parent.Left == node {
			return LeftSide, nil
		} else {
			return RightSide, nil
		}
	}
	return None, fmt.Errorf("try to define the parent's side for nil node")
}

func (tree *Tree) min(node *Node) (*Node, error) {
	for curNode := node; curNode != nil; curNode = curNode.Left {
		if curNode.Left == nil {
			return curNode, nil
		}
	}
	return nil, fmt.Errorf("try to find the min element in the empty binary tree")
}

func (tree *Tree) max(node *Node) (*Node, error) {
	for curNode := node; curNode != nil; curNode = curNode.Right {
		if curNode.Right == nil {
			return curNode, nil
		}
	}
	return nil, fmt.Errorf("try to find the max element in the empty binary tree")
}

func (tree *Tree) Add(key int64, val string) error {
	var newNode = &Node{
		Key:   key,
		Value: val,
	}

	if tree.Root == nil {
		tree.Root = newNode
	} else {
		var parent *Node

		for curNode := tree.Root; curNode != nil; {
			parent = curNode
			if curNode.Key > key {
				curNode = curNode.Left
			} else if curNode.Key < key {
				curNode = curNode.Right
			} else {
				return fmt.Errorf("try to add the existing node with key %v", key)
			}
		}

		newNode.Parent = parent

		if parent.Key > key {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}
	}

	return nil
}

func (tree *Tree) Set(key int64, val string) error {
	if node := tree.find(key); node != nil {
		node.Value = val
		return nil
	}
	return fmt.Errorf("try to set new value to unexisting node with key %v", key)
}

func (tree *Tree) Search(key int64) (string, error) {
	if node := tree.find(key); node != nil {
		return node.Value, nil
	}
	return "", fmt.Errorf("try to find the unexisting node with key %v", key)
}

func (tree *Tree) Min() (int64, string, error) {
	minNode, err := tree.min(tree.Root)
	return minNode.Key, minNode.Value, err
}

func (tree *Tree) Max() (int64, string, error) {
	maxNode, err := tree.max(tree.Root)
	return maxNode.Key, maxNode.Value, err
}

func (tree *Tree) delete(node *Node) error {
	if node != nil {
		side, _ := tree.defineParentSide(node)

		if node.Left == nil && node.Right == nil {
			if side == Root {
				tree.Root = nil
			} else if side == LeftSide {
				node.Parent.Left = nil
			} else if side == RightSide {
				node.Parent.Right = nil
			}
		} else if node.Right != nil && node.Left == nil {
			if side == Root {
				tree.Root = node.Right
			} else if side == LeftSide {
				node.Parent.Left = node.Right
			} else if side == RightSide {
				node.Parent.Right = node.Right
			}
			node.Right.Parent = node.Parent
		} else if node.Left != nil && node.Right == nil {
			if side == Root {
				tree.Root = node.Left
			} else if side == LeftSide {
				node.Parent.Left = node.Left
			} else if side == RightSide {
				node.Parent.Right = node.Left
			}
			node.Left.Parent = node.Parent
		}
		return nil
	}
	return fmt.Errorf("try to delete the unexisting node")
}

func (tree *Tree) Delete(key int64) error {
	if node := tree.find(key); node != nil {
		if node.Right == nil || node.Left == nil {
			tree.delete(node)
		} else {
			maxNode, _ := tree.max(node.Left)
			node.Key = maxNode.Key
			node.Value = maxNode.Value
			tree.delete(maxNode)
		}
		return nil
	}

	return fmt.Errorf("try to delete the unexisting node with key %v", key)
}

func (tree *Tree) print(w io.Writer) {
	var checkVert = false
	var buffer = make([]string, 0, 200)
	var queue = make([]*Node, 0, 200)
	queue = append(queue, tree.Root)

	for vertCount, lvlNum := 0, 1; len(queue) != 0; vertCount += 2 {
		if int(math.Pow(2, float64(lvlNum))) == vertCount {
			if !checkVert {
				break
			}
			for _, elems := range buffer {
				fmt.Fprint(w, elems)
			}
			fmt.Fprintln(w)
			vertCount = 0
			lvlNum++
			checkVert = false
			buffer = buffer[:0]
		}

		curNode := queue[0]
		queue = queue[1:]

		if curNode == nil || curNode.Left == nil {
			buffer = append(buffer, "_ ")
			queue = append(queue, (*Node)(nil))
		} else {
			buffer = append(buffer, fmt.Sprintf("[%v %v %v] ", curNode.Left.Key, curNode.Left.Value, curNode.Key))
			queue = append(queue, curNode.Left)
			checkVert = true
		}

		if curNode == nil || curNode.Right == nil {
			buffer = append(buffer, "_ ")
			queue = append(queue, (*Node)(nil))
		} else {
			buffer = append(buffer, fmt.Sprintf("[%v %v %v] ", curNode.Right.Key, curNode.Right.Value, curNode.Key))
			queue = append(queue, curNode.Right)
			checkVert = true
		}
	}
}

func (tree *Tree) Print(w io.Writer) {
	if tree.Root != nil {
		fmt.Fprintf(w, "[%v %v]\n", tree.Root.Key, tree.Root.Value)
		tree.print(w)
	} else {
		fmt.Fprintln(w, "_")
	}
}

func parseInput(input string) ([]string, error) {
	var commds = strings.Split(input, " ")

	if len(commds) != 0 && (((commds[0] == "add" || commds[0] == "set") && len(commds) == 3) ||
		((commds[0] == "delete" || commds[0] == "search") && len(commds) == 2) ||
		((commds[0] == "print" || commds[0] == "min" || commds[0] == "max") &&
			len(commds) == 1)) {
		return commds, nil
	} else {
		return nil, fmt.Errorf("error")
	}
}

func main() {
	var tree = NewTree()
	var treeOps = make([][]string, 0, 100)
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()

		if input == "" {
			continue
		}
		commds, err := parseInput(input)

		if err != nil {
			treeOps = append(treeOps, []string{"error"})
		} else {
			treeOps = append(treeOps, commds)
		}
	}

	for _, ops := range treeOps {
		if ops[0] == "add" {
			key, _ := strconv.Atoi(ops[1])
			err := tree.Add(int64(key), ops[2])

			if err != nil {
				fmt.Println("error")
			}
		} else if ops[0] == "set" {
			key, _ := strconv.Atoi(ops[1])
			err := tree.Set(int64(key), ops[2])

			if err != nil {
				fmt.Println("error")
			}
		} else if ops[0] == "delete" {
			key, _ := strconv.Atoi(ops[1])
			err := tree.Delete(int64(key))

			if err != nil {
				fmt.Println("error")
			}
		} else if ops[0] == "search" {
			key, _ := strconv.Atoi(ops[1])
			val, err := tree.Search(int64(key))

			if err != nil {
				fmt.Println("0")
			} else {
				fmt.Printf("1 %v\n", val)
			}
		} else if ops[0] == "print" {
			tree.Print(os.Stdout)

		} else if ops[0] == "min" {
			key, val, err := tree.Min()

			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Printf("%v %v\n", key, val)
			}
		} else if ops[0] == "max" {
			key, val, err := tree.Max()

			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Printf("%v %v\n", key, val)
			}
		} else {
			fmt.Println("error")
		}
	}
}
