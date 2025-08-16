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

// divideNode defines the logic of dividing the node by two
// nodes with amount of elements = Btree.factor - 1.
func (b *Btree[KeyType]) divideNode(node *node[KeyType]) {

}

// Add defines the logic of adding the key into the btree with the current key and val.
func (b *Btree[KeyType]) Add(key KeyType, val any) {
	node := b.root
	for !node.isLeaf() {
		for i := 0; i != len(node.childs); i++ {
			if i == 0 && node.childs[i].data.key > key ||
				(i != 0 && node.childs[i-1].data.key <= key && node.childs[i].data.key < key) ||
				node.childs[i].flagRightTreePtr {
				node = node.childs[i].ptr
				break
			}
		}
	}

	idx, flagExist := slices.BinarySearchFunc(node.childs, child[KeyType]{
		data: data[KeyType]{
			key: key,
		},
	}, compareChildsToIncrease)

	_ = idx
	_ = flagExist

	if len(node.childs) == b.factor*2-1 {
		// TODO: add here the dividing leaf procedure.
	}
	// TODO: add here the key inserting instructions':
	// 1. Add the element to the second slice if the start slice was divided.
	// 2. Add the element to the start slice if it wasn't divided.
}

func (b *Btree[KeyType]) Delete(key KeyType) {

}

func (b *Btree[KeyType]) Print(reciever io.Writer) {

}
