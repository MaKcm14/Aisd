package btree

// KeyComparable defines the type's interface describes the possible key types.
type KeyComparable interface {
	int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 |
		int | uint | string
}

// data defines the btree's data pair.
type data[KeyType KeyComparable] struct {
	key KeyType
	val any
}

// child defines the children description of the btree's node.
type child[KeyType KeyComparable] struct {
	data data[KeyType]
	ptr  *node[KeyType]
}

// node defines the node of the btree.
type node[KeyType KeyComparable] struct {
	childs []child[KeyType]
}

func newNode[KeyType KeyComparable]() *node[KeyType] {
	return &node[KeyType]{
		childs: make([]child[KeyType], 0),
	}
}
