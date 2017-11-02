package skiplist

import (
	"fmt"
	"math/rand"
)

type SkipList struct {
	Head *Node
}

type Node struct {
	Pointers []*Node
	Key      int
	Val      interface{}
}

func New(maxHeight int) *SkipList {
	return &SkipList{
		Head: &Node{
			Pointers: make([]*Node, maxHeight),
		},
	}
}

func (s *SkipList) Put(key int, val interface{}) *Node {
	newNodeHeight := randHeight(len(s.Head.Pointers))

	newNode := &Node{
		Pointers: make([]*Node, newNodeHeight),
		Key:      key,
		Val:      val,
	}

	node := s.Head
	for i := len(s.Head.Pointers); i > 0; i-- {
		indexNew := len(newNode.Pointers) - i
		indexTest := len(node.Pointers) - i
		nodeTest := node.Pointers[indexTest]

		for nodeTest != nil && nodeTest.Key <= key {
			node = node.Pointers[indexTest]
			indexTest = len(node.Pointers) - i
			nodeTest = node.Pointers[indexTest]
		}

		if i <= newNodeHeight {
			newNode.Pointers[indexNew] = node.Pointers[indexTest]
			node.Pointers[indexTest] = newNode
		}
	}

	return newNode
}

func (s *SkipList) Remove(toRemove *Node) bool {
	val := toRemove.Key
	removed := false
	node := s.Head
	for i := len(s.Head.Pointers); i > 0; i-- {
		indexTest := len(node.Pointers) - i
		nodeTest := node.Pointers[indexTest]

		for nodeTest != nil && nodeTest.Key < val {
			node = nodeTest
			indexTest = len(node.Pointers) - i
			nodeTest = node.Pointers[indexTest]
		}

		// update pointer while handling duplicates
		if nodeTest != nil && nodeTest.Key == val {
			dupNode := node
			dupNodeTest := nodeTest
			for dupNodeTest != nil && dupNodeTest.Key == val && dupNodeTest != toRemove {
				dupNode = dupNodeTest
				indexTest = len(dupNode.Pointers) - i
				dupNodeTest = dupNode.Pointers[indexTest]
			}

			if dupNodeTest == toRemove {
				removedNext := len(toRemove.Pointers) - i
				dupNode.Pointers[indexTest] = toRemove.Pointers[removedNext]
				removed = true
			}
		}
	}

	return removed
}

func (s *SkipList) Get(val int) (*Node, bool) {
	node := s.Head
	for i := len(s.Head.Pointers); i > 0; i-- {
		indexTest := len(node.Pointers) - i
		nodeTest := node.Pointers[indexTest]

		for nodeTest != nil && nodeTest.Key <= val {
			node = nodeTest
			indexTest = len(node.Pointers) - i
			nodeTest = node.Pointers[indexTest]
		}
	}

	return node, node != nil
}

func (s *SkipList) Dump() {
	for i := len(s.Head.Pointers) - 1; i >= 0; i-- {
		vals := []int{}
		for _, v := range s.ToSliceAtHeight(i) {
			vals = append(vals, v.Key)
		}
		fmt.Printf("Height: %d, Count: %d, %v\n", i, len(vals), vals)
	}
}

func (s *SkipList) ToSlice() []*Node {
	return s.ToSliceAtHeight(0)
}

func (s *SkipList) ToSliceAtHeight(height int) []*Node {
	res := []*Node{}
	node := s.Head.Pointers[len(s.Head.Pointers)-1-height]
	for node != nil {
		res = append(res, node)
		node = node.Pointers[len(node.Pointers)-1-height]
	}
	return res
}

func randHeight(maxHeight int) int {
	h := 1
	for ; rand.Intn(2) == 0 && h < maxHeight; h++ {
	}

	return h
}
