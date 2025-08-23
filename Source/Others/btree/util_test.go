package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareChildsToIncrease_PositiveCases(t *testing.T) {
	t.Run("Test_Elem_Child_Matches_Target", func(t *testing.T) {
		elem := child[int]{
			data: data[int]{
				key: 0,
			},
		}
		target := child[int]{
			data: data[int]{
				key: 0,
			},
		}
		assert.Equal(t, 0, compareChildsToIncrease(elem, target),
			"unexpected comparing: two equal values are not equal by comparing the func")
	})

	t.Run("Test_Elem_Greater_Child_Target", func(t *testing.T) {
		elem := child[int]{
			data: data[int]{
				key: 100,
			},
		}
		target := child[int]{
			data: data[int]{
				key: 0,
			},
		}
		assert.Positive(t, compareChildsToIncrease(elem, target),
			"unexpected comparing: values that must be in the order 'target, elem' "+
				" actually are in the other order")
	})

	t.Run("Test_Elem_Less_Child_Target", func(t *testing.T) {
		elem := child[int]{
			data: data[int]{
				key: -100,
			},
		}
		target := child[int]{
			data: data[int]{
				key: 0,
			},
		}
		assert.Negative(t, compareChildsToIncrease(elem, target),
			"unexpected comparing: values that must be in the order 'elem, target' "+
				" actually are in the other order")
	})
}

func TestInsert_PositiveCases(t *testing.T) {
	t.Run("Test_One_Elem_End_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3}
		insert(&slice, len(slice), 111)

		assert.Equal(t, []int{1, 2, 3, 111}, slice,
			"unexpected slice value: after inserting one element "+
				"at the end another value was got",
		)
	})

	t.Run("Test_Few_Elem_End_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3}
		insert(&slice, len(slice), 111, 112, 113, 114)

		assert.Equal(t, []int{1, 2, 3, 111, 112, 113, 114}, slice,
			"unexpected slice value: after inserting few values at the end of the slice "+
				"another slice's value was got",
		)
	})

	t.Run("Test_One_Elem_Internal_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		insert(&slice, len(slice)/2, 111)

		assert.Equal(t, []int{1, 2, 111, 3, 4, 5}, slice,
			"unexpected slice value: after inserting one element at the mid of the slice "+
				"another slice's value was got",
		)
	})

	t.Run("Test_Few_Elem_Internal_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		insert(&slice, len(slice)/2, 111, 112, 113, 114)

		assert.Equal(t, []int{1, 2, 111, 112, 113, 114, 3, 4, 5}, slice,
			"unexpected slice value: after inserting few elements at the mid of the slice "+
				"another slice's value was got",
		)
	})
}

func TestInsert_CornerCases(t *testing.T) {
	t.Run("Test_Empty_Value_Slice_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		insert(&slice, len(slice)/2)

		assert.Equal(t, []int{1, 2, 3, 4, 5}, slice,
			"unexpected slice value: after inserting an empty slice elems "+
				"at the mid of the slice another slice value was got",
		)

		slice = []int{1, 2, 3, 4, 5}
		insert(&slice, len(slice))

		assert.Equal(t, []int{1, 2, 3, 4, 5}, slice,
			"unexpected slice value: after inserting an ampty slice elems "+
				"at the end of the slice another slice value was got",
		)
	})
}

func TestInsert_NegativeCases(t *testing.T) {
	t.Run("Test_Wrong_Pos_Inserting", func(t *testing.T) {
		slice := []int{1, 2, 3}
		insert(&slice, -1, 0)

		assert.Equal(t, []int{1, 2, 3}, slice,
			"unexpected slice value: the internal array has been changed",
		)

		slice = []int{1, 2, 3}
		insert(&slice, len(slice)+1, 0)

		assert.Equal(t, []int{1, 2, 3}, slice,
			"unexpected slice value: the internal array has been changed",
		)
	})
}
