package btree

import "io"

// Btree is the main structure of btree data structure.
type Btree[KeyType KeyComparable] struct {
	factor int
	root   *node[KeyType]
}

func New[KeyType KeyComparable](factor int) Btree[KeyType] {
	return Btree[KeyType]{
		factor: factor,
		root:   newNode[KeyType](),
	}
}

// Find searches and returns the existing flag and data value for the current key.
func (b *Btree[KeyType]) Find(key KeyType) (bool, any) {
	for node := b.root; len(node.childs) != 0; {
		for i := 0; i != len(node.childs); i++ {
			if node.childs[i].data.key == key {
				return true, node.childs[i].data.val
			}
			if i == 0 && node.childs[i].data.key > key ||
				i == len(node.childs)-1 && node.childs[i].data.key < key ||
				(i != 0 && node.childs[i-1].data.key < key && node.childs[i+1].data.key < key) {
				node = node.childs[i].ptr
				break
			}
		}
	}
	return false, nil
}

func (b *Btree[KeyType]) Add(key KeyType, val any) {

}

func (b *Btree[KeyType]) Delete(key KeyType) {

}

func (b *Btree[KeyType]) Print(reciever io.Writer) {

}
