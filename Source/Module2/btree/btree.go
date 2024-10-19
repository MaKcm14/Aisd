package btree

import (
	"fmt"
	"math"
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

func (tree *Tree) successor(node *Node) (*Node, error) {
	if node == nil {
		return nil, fmt.Errorf("try to find the successor element for the empty node")
	}

	if node.Right != nil {
		return tree.min(node.Right)
	}

	var successor = node.Parent
	for successor != nil && node == successor.Right {
		node = successor
		successor = successor.Parent
	}
	return successor, nil
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

func (tree *Tree) Delete(key int64) error {
	if node := tree.find(key); node != nil {
		side, _ := tree.defineParentSide(node)

		if node.Left == nil && node.Right == nil {
			if side == Root {
				tree.Root = nil
			} else if side == LeftSide {
				node.Parent.Left = nil
			} else if side == RightSide {
				node.Parent.Right = nil
			}
		} else if node.Right != nil {
			if side == Root {
				tree.Root = node.Right
			} else if side == LeftSide {
				node.Parent.Left = node.Right
			} else if side == RightSide {
				node.Parent.Right = node.Right
			}
		} else if node.Left != nil {
			if side == Root {
				tree.Root = node.Left
			} else if side == LeftSide {
				node.Parent.Left = node.Left
			} else if side == RightSide {
				node.Parent.Right = node.Left
			}
		} else {
			successor, _ := tree.successor(node)
			sucSide, _ := tree.defineParentSide(successor)
			minSuc, _ := tree.min(successor)

			if side != Root {
				if sucSide == LeftSide {
					successor.Parent.Left = nil
				} else if sucSide == RightSide {
					successor.Parent.Right = nil
				}
				successor.Parent = node.Parent
				if side == LeftSide {
					successor.Parent.Left = successor
				} else if side == RightSide {
					successor.Parent.Right = successor
				}
				minSuc.Left = node.Left
				node.Left.Parent = minSuc
			} else {
				tree.Root.Right.Parent = successor
				if sucSide == LeftSide {
					successor.Parent.Left = nil
				} else if sucSide == RightSide {
					successor.Parent.Right = nil
				}
				successor.Parent = nil
				minSuc.Left = tree.Root.Left
				tree.Root.Left.Parent = minSuc
				tree.Root = successor
			}
		}

		return nil
	}

	return fmt.Errorf("try to delete the unexisting node with key %v", key)
}

func (tree *Tree) print() {
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
				fmt.Print(elems)
			}
			fmt.Print("\n")
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

func (tree *Tree) Print() {
	if tree.Root != nil {
		fmt.Printf("[%v %v]\n", tree.Root.Key, tree.Root.Value)
		tree.print()
	} else {
		fmt.Println("_")
	}
}
