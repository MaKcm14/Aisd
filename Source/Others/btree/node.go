package btree

// KeyInt defines the integer interface for using in the KeyType.
type KeyInt interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8
}

// KeyUInt defines the non-negative integer interface for using in the KeyType.
type KeyUInt interface {
	~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8
}

// KeyComparable defines the type's interface describes the possible key types.
// See the KeyInt and KeyUInt for more info about possible types.
type KeyComparable interface {
	KeyInt | KeyUInt | ~string
}

// data defines the btree's data pair and stores the supremum of every key in its child node.
type data[KeyType KeyComparable] struct {
	key KeyType
	val any
}

// child defines the children description of the btree's node.
type child[KeyType KeyComparable] struct {
	data             data[KeyType]
	ptr              *node[KeyType]
	flagRightTreePtr bool
}

// node defines the node of the btree.
type node[KeyType KeyComparable] struct {
	// childs defines the every node's child.
	childs  []child[KeyType]
	pParent *node[KeyType]
}

func newNode[KeyType KeyComparable](factor int) *node[KeyType] {
	return &node[KeyType]{
		childs: make([]child[KeyType], 0, 2*factor),
	}
}

// isLeaf defines whether the current node is the leaf.
func (n node[KeyType]) isLeaf() bool {
	for _, child := range n.childs {
		if child.ptr != nil {
			return false
		}
	}
	return true
}

// updayeChildsParent defines the logic of updating the child's parent ptr to the current node.
func (n *node[KeyType]) updateChildsParent(pParent *node[KeyType]) {
	for _, child := range n.childs {
		if child.ptr != nil {
			child.ptr.pParent = pParent
		}
	}
}

// getKeyAmount returns the amount of the node's key.
func (n *node[KeyType]) getKeyAmount() int {
	if l := len(n.childs); l != 0 && n.childs[l-1].flagRightTreePtr {
		return l - 1
	} else if l == 0 {
		return 0
	}
	return len(n.childs)
}
