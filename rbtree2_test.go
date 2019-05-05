package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Balanced check if redblack tree check if
// path from the root node to any of its descendant NIL nodes contains the same number of black nodes.
func (t *Tree) IsBalanced() bool {
	if t.root == nil {
		return true
	}
	blackNums := 0
	fisrt := 0
	return t.root.blackNodes(blackNums, &fisrt)
}

// first is used to record the number of black nodes in first path from node n
// to its descendant NIL nodes.
func (n *node) blackNodes(blacks int, first *int) bool {
	if n.color == BLACK {
		blacks++
	}
	if n.left != nil {
		if !n.left.blackNodes(blacks, first) {
			return false
		}
	} else {
		if *first == 0 {
			*first = blacks
		} else if *first != blacks {
			return false
		}
	}
	if n.right != nil {
		if !n.right.blackNodes(blacks, first) {
			return false
		}
	} else {
		if *first == 0 {
			*first = blacks
		} else if blacks != *first {
			return false
		}
	}
	return true
}

func TestRandomInsertAndDelete(t *testing.T) {
	ltree := NewTree()
	var arr []int
	dict := map[int]bool{}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10000; i++ {
		lrand := rand.Intn(30)
		if lrand > 25 {
			if len(arr) == 0 {
				continue
			}
			idx := rand.Intn(len(arr))
			v := arr[idx]
			arr = remove(arr, idx)
			delete(dict, v)
			ltree.Delete(key(v))
			if !ltree.IsBalanced() {
				ltree.Preorder()
				panic(fmt.Sprintf("after %d deleted, tree isn't balanced", v))
			}
			if len(arr) != ltree.Size() {
				t.Errorf("tree's size is not right, expect %d, got %d", len(arr), ltree.Size())
			}
		} else {
			v := rand.Intn(100000)
			if dict[v] {
				// duplicate key
				continue
			}
			arr = append(arr, v)
			dict[v] = true
			ltree.Insert(key(v), v)
			if !ltree.IsBalanced() {
				ltree.Preorder()
				panic(fmt.Sprintf("after inserting %d, tree isn't balanced", v))
			}
		}
	}

	// random delete
	for i := 0; i < 10000; i++ {
		v := rand.Intn(1000000)
		ltree.Delete(key(v))
		if !ltree.IsBalanced() {
			ltree.Preorder()
			panic(fmt.Sprintf("after %d deleted, tree isnot balance", v))
		}
	}
}

func remove(arr []int, idx int) []int {
	if idx >= len(arr) {
		return arr
	}
	if idx == len(arr)-1 {
		return arr[:idx]
	}
	return append(arr[:idx], arr[idx+1:]...)
}
