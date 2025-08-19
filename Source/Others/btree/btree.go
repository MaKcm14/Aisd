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
	for node := b.root; !node.isLeaf(); {
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

// splitNode defines the logic of dividing the node by two
// nodes when the node's amount of elements = 2 * Btree.factor - 1.
// It returns two pointers on two parts of the divided node.
func (b *Btree[KeyType]) splitNode(node *node[KeyType]) (*node[KeyType], *node[KeyType]) {
	if len(node.childs) < 2*b.factor-1 {
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
	nodeLeftRightTree := nodeMedian
	nodeLeftRightTree.flagRightTreePtr = true
	nodeLeftRightTree.data = data[KeyType]{}

	nodeLeft.childs = append(nodeLeft.childs, nodeLeftRightTree)

	if node.pParent != nil {
		if len(node.pParent.childs) < 2*b.factor-1 {
			idx, _ := slices.BinarySearchFunc(node.pParent.childs,
				nodeMedian, compareChildsToIncrease,
			)
			nodeMedian.ptr = nodeLeft

			insert(&node.pParent.childs, idx, nodeMedian)

			nodeLeft.pParent = node.pParent
			nodeRight.pParent = node.pParent

			nodeLeft.updateChildsParent()
			nodeRight.updateChildsParent()

			// Fix the pointers ptr from the childs of the nodeLeft.pParent.
			for _, child := range nodeLeft.pParent.childs {
				if child.ptr == node {
					child.ptr = nodeLeft
				}
			}
			nodeLeftParentChild := node.pParent.childs[idx]
			nodeLeftParentChild.ptr = nodeLeft
			node.pParent.childs[idx] = nodeLeftParentChild

			nodeRightParentChild := node.pParent.childs[idx+1]
			nodeRightParentChild.ptr = nodeRight
			node.pParent.childs[idx+1] = nodeRightParentChild

			return nodeLeft, nodeRight

		} else {
			pParentLeft, pParentRight := b.splitNode(node.pParent)
			_, _ = pParentLeft, pParentRight

			// TODO: add the logic of dividing the parent and spliting the current node.
			// Don't forget about the pointers rereference.
		}
	} else {
		// TODO: add the logic of creating a new Btree's root node and make rereference of
		// the current node's parts.
	}

	return nil, nil
}

// Add defines the logic of adding the key into the btree with the current key and val.
func (b *Btree[KeyType]) Add(key KeyType, val any) {
	// TODO: define here the logic of adding the element to the btree:
	// lookup the needed node (leaf), divide if need and then add the element.
}

func (b *Btree[KeyType]) Delete(key KeyType) {

}

func (b *Btree[KeyType]) Print(reciever io.Writer) {

}
