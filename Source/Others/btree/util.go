package btree

// compareChildsToIncrease defines the order to compare the elements of the childs' slice
// in order to increasing the elements.
func compareChildsToIncrease[KeyType KeyComparable](elem child[KeyType], target child[KeyType]) int {
	if elem.data.key == target.data.key && !elem.flagRightTreePtr {
		return 0
	} else if elem.data.key < target.data.key && !elem.flagRightTreePtr {
		return -1
	}
	return 1
}

// insert realises the logic of inserting the values into the slice starting the pos.
func insert[Type any](slice *[]Type, pos int, vals ...Type) {
	lastSecondChilds := append(
		make([]Type, 0, len(*slice)),
		(*slice)[pos:]...,
	)
	*slice = (*slice)[:pos]
	*slice = append(*slice, vals...)
	*slice = append(*slice, lastSecondChilds...)
}
