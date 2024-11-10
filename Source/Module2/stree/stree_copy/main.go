package stree_copy

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	startVert := tree.Root
	for startVert != nil {
		if startVert.Key > key {
			if startVert.Left == nil {
				tree.splay(startVert)
				break
			}
			startVert = startVert.Left
		} else if startVert.Key < key {
			if startVert.Right == nil {
				tree.splay(startVert)
				break
			}
			startVert = startVert.Right
		} else {
			tree.splay(startVert)
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

func (tree *Tree) leftRotate(node *Node) error {
	if node == nil || node.Right == nil {
		return errors.New("try to rotate in the left side the empty node (is nil)")
	}
	nodeSide, _ := tree.defineParentSide(node)
	changeNode := node.Right

	changeNode.Parent = node.Parent
	if nodeSide == Root {
		tree.Root = changeNode
	} else if nodeSide == LeftSide {
		node.Parent.Left = changeNode
	} else {
		node.Parent.Right = changeNode
	}

	node.Parent = changeNode
	node.Right = changeNode.Left
	if changeNode.Left != nil {
		changeNode.Left.Parent = node
	}
	changeNode.Left = node

	return nil
}

func (tree *Tree) rightRotate(node *Node) error {
	if node == nil || node.Left == nil {
		return errors.New("try to rotate in the right side the empty node (is nil)")
	}
	nodeSide, _ := tree.defineParentSide(node)
	changeNode := node.Left

	changeNode.Parent = node.Parent
	if nodeSide == Root {
		tree.Root = changeNode
	} else if nodeSide == LeftSide {
		node.Parent.Left = changeNode
	} else {
		node.Parent.Right = changeNode
	}

	node.Parent = changeNode
	node.Left = changeNode.Right
	if changeNode.Right != nil {
		changeNode.Right.Parent = node
	}
	changeNode.Right = node

	return nil
}

func (tree *Tree) splay(node *Node) error {
	if node == nil {
		return errors.New("try to splay the empty vertex (is nil)")
	}

	for node != tree.Root {
		nodeSide, _ := tree.defineParentSide(node)

		if nodeSide == LeftSide && node.Parent == tree.Root {
			tree.rightRotate(node.Parent)
			break
		} else if nodeSide == RightSide && node.Parent == tree.Root {
			tree.leftRotate(node.Parent)
			break
		} else if nodeSide != None {
			parentSide, _ := tree.defineParentSide(node.Parent)

			if parentSide == LeftSide && nodeSide == LeftSide {
				tree.rightRotate(node.Parent.Parent)
				tree.rightRotate(node.Parent)
			} else if parentSide == RightSide && nodeSide == RightSide {
				tree.leftRotate(node.Parent.Parent)
				tree.leftRotate(node.Parent)
			} else if parentSide == LeftSide && nodeSide == RightSide {
				tree.leftRotate(node.Parent)
				tree.rightRotate(node.Parent)
			} else if parentSide == RightSide && nodeSide == LeftSide {
				tree.rightRotate(node.Parent)
				tree.leftRotate(node.Parent)
			}
		}
	}

	return nil
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
				tree.splay(curNode)
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

	tree.splay(newNode)

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

	if minNode == nil {
		return 0, "", err
	}

	tree.splay(minNode)
	return minNode.Key, minNode.Value, err
}

func (tree *Tree) Max() (int64, string, error) {
	maxNode, err := tree.max(tree.Root)

	if maxNode == nil {
		return 0, "", err
	}

	tree.splay(maxNode)
	return maxNode.Key, maxNode.Value, err
}

func (tree *Tree) Delete(key int64) error {
	if node := tree.find(key); node != nil {
		if node.Left != nil && node.Right != nil {
			leftTree := Tree{tree.Root.Left}
			rightTree := Tree{tree.Root.Right}

			leftTree.Root.Parent = nil

			rightTree.Root.Parent = nil

			tree.Root = nil

			maxNode, _ := leftTree.max(leftTree.Root)

			leftTree.splay(maxNode)

			tree.Root = leftTree.Root

			tree.Root.Right = rightTree.Root
			rightTree.Root.Parent = tree.Root

		} else if node.Left == nil && node.Right == nil {
			tree.Root = nil
		} else if node.Left == nil || node.Right == nil {
			if node.Left == nil {
				tree.Root = node.Right
			} else if node.Right == nil {
				tree.Root = node.Left
			}
			tree.Root.Parent = nil
		}
		return nil
	}
	return fmt.Errorf("try to delete the unexisting node with key %v", key)
}

func (tree *Tree) Print() {
	if tree.Root == nil {
		fmt.Println("_")
		return
	}

	var curLvl = []*Node{tree.Root}
	var lvlNum = 0

	fmt.Printf("[%d %s]", tree.Root.Key, tree.Root.Value)

	for checkVertFlag := false; len(curLvl) != 0 && !checkVertFlag; {
		var nextLvl = make([]*Node, 0, cap(curLvl)*2)

		checkVertFlag = true
		for _, node := range curLvl {
			if node != nil {
				if lvlNum != 0 {
					fmt.Printf("[%d %s %s] ", node.Key, node.Value, fmt.Sprintf("%d", node.Parent.Key))
				}
				nextLvl = append(nextLvl, node.Left, node.Right)
				if node.Left != nil || node.Right != nil {
					checkVertFlag = false
				}
			} else {
				fmt.Print("_ ")
				nextLvl = append(nextLvl, nil, nil)
			}
		}
		lvlNum = 1
		fmt.Print("\n")

		curLvl = nextLvl
	}
}

func (tree *Tree) TestPrint() {
	if tree.Root == nil {
		fmt.Println("_")
		return
	}
	curLen, curLayer, nextLayer, stop := 1, map[int]*Node{0: tree.Root}, map[int]*Node{}, false
	for !stop {
		stop = true
		for i := 0; i < curLen; i++ {
			node := curLayer[i]
			if node == nil {
				if i == 0 {
					fmt.Print("_")
				} else {
					fmt.Print(" _")
				}
				continue
			}
			if node.Left != nil {
				stop = false
				nextLayer[i*2] = node.Left
			}
			if node.Right != nil {
				stop = false
				nextLayer[i*2+1] = node.Right
			}
			if node.Parent == nil {
				fmt.Printf("[%d %s]", node.Key, node.Value)
			} else if i == 0 {
				fmt.Printf("[%d %s %d]", node.Key, node.Value, node.Parent.Key)
			} else {
				fmt.Printf(" [%d %s %d]", node.Key, node.Value, node.Parent.Key)
			}
		}
		curLayer, nextLayer = nextLayer, map[int]*Node{}
		curLen *= 2
		fmt.Println()
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
	var slice = [][]float32{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}

	var tree = NewTree()
	var scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var input = scanner.Text()

		if input == "" {
			continue
		}
		ops, err := parseInput(input)

		if err != nil {
			ops = []string{"error"}
		}

		if ops[0] == "add" {
			var timerAdd = time.Now()

			key, _ := strconv.Atoi(ops[1])
			err := tree.Add(int64(key), ops[2])

			if err != nil {
				fmt.Println("error")
			}
			slice[0][0] += float32(time.Since(timerAdd).Seconds())
			slice[0][1]++

		} else if ops[0] == "set" {
			var timerSet = time.Now()

			key, _ := strconv.Atoi(ops[1])
			err := tree.Set(int64(key), ops[2])

			if err != nil {
				fmt.Println("error")
			}
			slice[1][0] += float32(time.Since(timerSet).Seconds())
			slice[1][1]++

		} else if ops[0] == "delete" {
			var timerDelete = time.Now()

			key, _ := strconv.Atoi(ops[1])
			err := tree.Delete(int64(key))

			if err != nil {
				fmt.Println("error")
			}

			slice[2][0] += float32(time.Since(timerDelete).Seconds())
			slice[2][1]++

		} else if ops[0] == "search" {
			var timerSearch = time.Now()

			key, _ := strconv.Atoi(ops[1])
			val, err := tree.Search(int64(key))

			if err != nil {
				fmt.Println("0")
			} else {
				fmt.Printf("1 %v\n", val)
			}
			slice[3][0] += float32(time.Since(timerSearch).Seconds())
			slice[3][1]++

		} else if ops[0] == "print" {
			var timerPrint = time.Now()

			tree.Print()

			slice[6][0] += float32(time.Since(timerPrint).Seconds())
			slice[6][1]++
		} else if ops[0] == "min" {
			var timerMin = time.Now()

			key, val, err := tree.Min()

			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Printf("%v %v\n", key, val)
			}

			slice[4][0] += float32(time.Since(timerMin).Seconds())
			slice[4][1]++

		} else if ops[0] == "max" {
			var timerMax = time.Now()

			key, val, err := tree.Max()

			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Printf("%v %v\n", key, val)
			}

			slice[5][0] += float32(time.Since(timerMax).Seconds())
			slice[5][1]++
		} else {
			fmt.Println("error")
		}
	}

	fmt.Print("\n\n")
	for index, timeRes := range slice {
		if index == 0 {
			fmt.Printf("Add: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 1 {
			fmt.Printf("Set: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 2 {
			fmt.Printf("Delete: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 3 {
			fmt.Printf("Search: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 4 {
			fmt.Printf("Min: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 5 {
			fmt.Printf("Max: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		} else if index == 6 {
			fmt.Printf("Print: %v %v\n", timeRes[0], timeRes[0]/timeRes[1])
		}
	}
}
