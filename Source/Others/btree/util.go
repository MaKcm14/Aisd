package btree

// compareChildsToIncrease defines the order to compare the elements of the childs' slice
// in order to increasing the elements:
// it returns 0 if the slice element matches the target,
// a negative number if the slice element precedes the target,
// or a positive number if the slice element follows the target
func compareChildsToIncrease[KeyType KeyComparable](elem child[KeyType], target child[KeyType]) int {
	if elem.data.key == target.data.key && !elem.flagRightTreePtr {
		return 0
	} else if elem.data.key < target.data.key && !elem.flagRightTreePtr {
		return -1
	}
	return 1
}

// insert realises the logic of inserting the values into the slice starting the pos.
// pos must be from range [0, len(slice)], where pos == len(slice) defines adding the
// vals into the end.
func insert[Type any](slice *[]Type, pos int, vals ...Type) {
	if pos < 0 || pos > len(*slice) {
		return
	}
	lastSecondChilds := append(
		make([]Type, 0, len(*slice)),
		(*slice)[pos:]...,
	)
	*slice = (*slice)[:pos]
	*slice = append(*slice, vals...)
	*slice = append(*slice, lastSecondChilds...)
}
