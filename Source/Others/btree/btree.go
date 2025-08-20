package btree

import (
	"io"
	"slices"
)

// Btree is the main structure of btree data structure.
type Btree[KeyType KeyComparable] struct {
	factor int
	root   *node[KeyType]
}

func New[KeyType KeyComparable](factor int) Btree[KeyType] {
	return Btree[KeyType]{
		factor: factor,
		root:   newNode[KeyType](factor),
	}
}

// Find searches and returns the existing flag and data value for the current key.
func (b *Btree[KeyType]) Find(key KeyType) (bool, any) {
	for node, lastNode := b.root, (*node[KeyType])(nil); node != nil && node != lastNode; {
		lastNode = node
		for i := 0; i != len(node.childs); i++ {
			if node.childs[i].data.key == key {
				return true, node.childs[i].data.val
			}
			if i == 0 && node.childs[i].data.key > key ||
				(i != 0 && node.childs[i-1].data.key <= key && node.childs[i].data.key < key) ||
				node.childs[i].flagRightTreePtr {
				node = node.childs[i].ptr
				break
			}
		}
	}
	return false, nil
}

// updateRoot defines the logic of changing the btree's root using the two divided
// last root's nodes nodeLeft and nodeRight and the nodeMedian from the last Btree.root.
func (b *Btree[KeyType]) updateRoot(nodeLeft, nodeRight *node[KeyType], nodeMedian child[KeyType]) {
	newRoot := newNode[KeyType](b.factor)

	nodeMedian.ptr = nodeLeft

	newRoot.childs = append(newRoot.childs, nodeMedian)
	newRoot.childs = append(newRoot.childs, child[KeyType]{
		ptr:              nodeRight,
		flagRightTreePtr: true,
	})

	b.root.updateChildsParent(newRoot)
	b.root = newRoot
}

// addMedianToParent defines the logic of adding the median of the divided node
// to the parent's childs.
func (b *Btree[KeyType]) addMedianToParent(nodeParent, nodeLeft, nodeRight *node[KeyType], nodeMedian child[KeyType]) {
	idx, _ := slices.BinarySearchFunc(nodeParent.childs, nodeMedian, compareChildsToIncrease)

	nodeMedian.ptr = nodeLeft
	insert(&nodeParent.childs, idx, nodeMedian)

	nodeLeft.pParent = nodeParent
	nodeRight.pParent = nodeParent

	nodeLeft.updateChildsParent(nodeLeft)
	nodeRight.updateChildsParent(nodeRight)
}

// splitNode defines the logic of dividing the node by two
// nodes when the node's amount of elements (keys) = 2 * Btree.factor - 1.
// It returns two pointers on two parts of the divided node.
func (b *Btree[KeyType]) splitNode(node *node[KeyType]) (*node[KeyType], *node[KeyType]) {
	if node.getKeyAmount() < 2*b.factor-1 {
		return nil, nil
	}

	var (
		midIdx    = b.factor - 1
		nodeLeft  = newNode[KeyType](b.factor)
		nodeRight = newNode[KeyType](b.factor)
	)

	nodeLeft.childs = append(nodeLeft.childs, node.childs[:midIdx]...)
	nodeRight.childs = append(nodeRight.childs, node.childs[midIdx+1:]...)

	nodeMedian := node.childs[midIdx]

	if nodeMedian.ptr != nil {
		nodeLeftRightTree := nodeMedian
		nodeLeftRightTree.flagRightTreePtr = true
		nodeLeftRightTree.data = data[KeyType]{}

		nodeLeft.childs = append(nodeLeft.childs, nodeLeftRightTree)
	}

	if node.pParent != nil {
		if node.pParent.getKeyAmount() < 2*b.factor-1 {
			b.addMedianToParent(node.pParent, nodeLeft, nodeRight, nodeMedian)
		} else {
			pParentLeft, pParentRight := b.splitNode(node.pParent)
			_, _ = pParentLeft, pParentRight

			// TODO: add the logic of dividing the parent and spliting the current node.
			// Don't forget about the pointers rereference.
		}
	} else {
		b.updateRoot(nodeLeft, nodeRight, nodeMedian)
	}

	return nodeLeft, nodeRight
}

// Add defines the logic of adding the key into the btree with the current key and val.
func (b *Btree[KeyType]) Add(key KeyType, val any) {
	node := b.root
	for node.isLeaf() {
		for i := 0; i != len(node.childs); i++ {
			if i == 0 && node.childs[i].data.key > key ||
				(i != 0 && node.childs[i-1].data.key <= key && node.childs[i].data.key < key) ||
				node.childs[i].flagRightTreePtr {
				node = node.childs[i].ptr
				break
			}
		}
	}

	pLeft, pRight := b.splitNode(node)

	if pLeft != nil && pRight != nil {
		// TODO: define the needed node (pLeft, pRight) and
		// add here adding the data{key, val} to the node's childs
		// one of two (pLeft, pRight).
	} else {
		// TODO: add here adding the data{key, val} to the node's childs
		// when there wasn't any split.
	}
}

func (b *Btree[KeyType]) Delete(key KeyType) {

}

func (b *Btree[KeyType]) Print(reciever io.Writer) {

}
