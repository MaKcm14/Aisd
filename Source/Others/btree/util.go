package btree

// compareChildsToIncrease defines the order to compare the elements of the childs' slice
// in order to increasing the elements.
func compareChildsToIncrease[KeyType KeyComparable](elem child[KeyType], target child[KeyType]) int {
	if elem.data.key == target.data.key {
		return 0
	} else if elem.data.key < target.data.key {
		return -1
	}
	return 1
}

// insert realises the logic of inserting the value element into the slice on the pos.
func insert[Type any](slice *[]Type, pos int, value Type) {
	if pos != len(*slice) {
		lastSecondChilds := append(
			make([]Type, 0, len(*slice)),
			(*slice)[pos:]...,
		)
		(*slice)[pos] = value
		*slice = append((*slice)[:pos+1], lastSecondChilds...)
	} else {
		(*slice) = append(*slice, value)
	}
}
