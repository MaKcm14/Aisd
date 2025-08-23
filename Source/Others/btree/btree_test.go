package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type splitNodeTestSuite struct {
	suite.Suite
	bTree Btree[int]
}

func (s *splitNodeTestSuite) setupBtreeRootNoRightSubTree() {
	const factor = 2
	s.bTree = New[int](factor) // 2 * 2 - 1 = 3 keys max; 3 + 1 = 4 childs available.

	s.bTree.root = newNode[int](factor)

	s.bTree.root.childs = []child[int]{
		{
			data: data[int]{
				key: 1,
			},
			ptr: &node[int]{
				childs: []child[int]{
					{
						data: data[int]{
							key: -10,
						},
					},
				},
				pParent: s.bTree.root,
			},
		},
		{
			data: data[int]{
				key: 5,
			},
			ptr: &node[int]{
				childs: []child[int]{
					{
						data: data[int]{
							key: 2,
						},
					},
				},
				pParent: s.bTree.root,
			},
		},
		{
			data: data[int]{
				key: 9,
			},
			ptr: &node[int]{
				childs: []child[int]{
					{
						data: data[int]{
							key: 6,
						},
					},
				},
				pParent: s.bTree.root,
			},
		},
	}
}

func (s *splitNodeTestSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestUpdateRoot_PositiveCase" {
		s.setupBtreeRootNoRightSubTree()
	} else if testName == "TestAddMedianToParent_PositiveCase" {

	}
}

func (s *splitNodeTestSuite) TestUpdateRoot_PositiveCase() {
	const factor = 2
	t := s.T()

	nodeLeft := newNode[int](factor)
	nodeRight := newNode[int](factor)

	nodeMedian := s.bTree.root.childs[1]

	assert.NotPanics(t, func() { s.bTree.updateRoot(nodeLeft, nodeRight, nodeMedian) },
		"unexpected panic: the updateRoot call mustn't panic with the correct state",
	)

	assert.Equal(t, 2, len(s.bTree.root.childs),
		"unexpected len of the root's childs",
	)
	assert.Equal(t, nodeLeft, s.bTree.root.childs[0].ptr,
		"unexpected left subtree config",
	)
	assert.Equal(t, nodeRight, s.bTree.root.childs[1].ptr,
		"unexpected right subtree config",
	)
	assert.Equal(t, s.bTree.root.childs[0].ptr.pParent, s.bTree.root,
		"unexpected left node config: the wrong parent's pointer was set",
	)
	assert.Equal(t, s.bTree.root.childs[1].ptr.pParent, s.bTree.root,
		"unexpected right node config: the wrong parent's pointer was set",
	)
	assert.Equal(t, true, s.bTree.root.childs[1].flagRightTreePtr,
		"unexpected Btree's config: the right subtree existing was waited",
	)
}

func (s *splitNodeTestSuite) TestAddMedianToParent_PositiveCase() {

}

func TestSplitNodeTestSuite(t *testing.T) {
	suite.Run(t, new(splitNodeTestSuite))
}
