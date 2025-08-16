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
