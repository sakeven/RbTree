package rbtree

import (
	"fmt"
	"strings"
	"testing"
)

// Preorder prints the tree in pre order
func (t *Tree[K, V]) Preorder() {
	fmt.Println("preorder begin!")
	if t.root != nil {
		t.root.preorder()
	}
	fmt.Println("preorder end!")
}

func (n *node[K, V]) preorder() {
	fmt.Printf("(%v %v) ", n.Key, n.Value)
	if n.parent == nil {
		fmt.Printf("nil")
	} else {
		fmt.Printf("whose parent is %v", n.parent.Key)
	}
	if n.color == RED {
		fmt.Println(" and color RED")
	} else {
		fmt.Println(" and color BLACK")
	}
	if n.left != nil {
		fmt.Printf("%v's left child is ", n.Key)
		n.left.preorder()
	}
	if n.right != nil {
		fmt.Printf("%v's right child is ", n.Key)
		n.right.preorder()
	}
}

func TestPreorder(t *testing.T) {
	tree := NewTree[int, string]()
	if !tree.Empty() {
		t.Error("tree not empty")
	}

	tree.Insert(1, "123")
	tree.Insert(3, "234")
	tree.Insert(4, "dfa3")
	tree.Insert(6, "sd4")
	tree.Insert(5, "jcd4")
	tree.Insert(2, "bcd4")
	if tree.Size() != 6 {
		t.Error("Error size")
	}
	if tree.Empty() {
		t.Error("tree empty")
	}
	tree.Preorder()
}

func TestTree(t *testing.T) {
	tree := NewTree[int, string]()

	t.Run("empty", func(t *testing.T) {
		tree.Clear()
		tree.Delete(1)

		if !tree.Empty() {
			t.Fatal("tree isn't empty")
		}

		size := tree.Size()
		if size != 0 {
			t.Errorf("unexpected tree size of %d", size)
		}
	})

	t.Run("nil", func(t *testing.T) {
		tree = nil

		tree.Clear()
		tree.Delete(1)

		if !tree.Empty() {
			t.Fatal("tree isn't empty")
		}

		size := tree.Size()
		if size != 0 {
			t.Errorf("unexpected tree size of %d", size)
		}
	})

	var caught interface{}
	t.Run("catch insert panic", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				caught = err
			}
		}()
		// panics
		tree.Insert(1, "abc")
	})

	error := fmt.Sprintf("%v", caught)
	if !strings.Contains(error, "nil pointer dereference") {
		t.Fatalf("unexpected error: %#v", caught)
	}
}

func TestFind(t *testing.T) {
	tree := NewTree[int, string]()

	tree.Insert(1, "123")
	tree.Insert(3, "234")
	tree.Insert(4, "dfa3")
	tree.Insert(6, "sd4")
	tree.Insert(5, "jcd4")
	tree.Insert(2, "bcd4")

	n := tree.FindIt(4)
	if n.Value != "dfa3" {
		t.Error("Error value")
	}
	n.Value = "bdsf"
	if n.Value != "bdsf" {
		t.Error("Error value modify")
	}
	value := tree.Find(5)
	if value != "jcd4" {
		t.Error("Error value after modifyed other node")
	}

	t.Run("empty", func(t *testing.T) {
		tree = NewTree[int, string]()

		n := tree.FindIt(4)
		if n != nil {
			t.Fatalf("got %#v", n)
		}

		value := tree.Find(5)
		if value != "" {
			t.Fatalf("got %q", value)
		}
	})

	t.Run("nil", func(t *testing.T) {
		tree = nil

		n := tree.FindIt(4)
		if n != nil {
			t.Fatalf("got %#v", n)
		}

		value := tree.Find(5)
		if value != "" {
			t.Fatalf("got %q", value)
		}
	})
}

func TestIterator(t *testing.T) {
	tree := NewTree[int, string]()

	tree.Insert(1, "123")
	tree.Insert(3, "234")
	tree.Insert(4, "dfa3")
	tree.Insert(6, "sd4")
	tree.Insert(5, "jcd4")
	tree.Insert(2, "bcd4")

	it := tree.Iterator()

	for it != nil {
		it = it.Next()
	}

	t.Run("empty", func(t *testing.T) {
		tree = NewTree[int, string]()

		next := tree.Iterator()
		if next != nil {
			t.Fatalf(".Iterator() returned %#v", next)
		}

		size := tree.Size()
		if size != 0 {
			t.Fatalf("got size %d", size)
		}
	})

	t.Run("nil", func(t *testing.T) {
		tree = nil

		next := tree.Iterator()
		if next != nil {
			t.Fatalf(".Iterator() returned %#v", next)
		}

		size := tree.Size()
		if size != 0 {
			t.Fatalf("got size %d", size)
		}
	})
}

func TestDelete(t *testing.T) {
	tree := NewTree[int, string]()

	tree.Insert(1, "123")
	tree.Insert(3, "234")
	tree.Insert(4, "dfa3")
	tree.Insert(6, "sd4")
	tree.Insert(5, "jcd4")
	tree.Insert(2, "bcd4")
	for i := 1; i <= 6; i++ {
		tree.Delete(i)
		if tree.Size() != 6-i {
			t.Error("Delete Error")
		}
	}
	tree.Insert(1, "bcd4")
	tree.Clear()
	tree.Preorder()
	if tree.Find(1) != "" {
		t.Error("Can't clear")
	}

	t.Run("empty", func(t *testing.T) {
		tree = NewTree[int, string]()
		tree.Delete(1)

		size := tree.Size()
		if size != 0 {
			t.Fatalf("after size is %d", size)
		}
	})

	t.Run("nil", func(t *testing.T) {
		tree = nil
		tree.Delete(1)

		size := tree.Size()
		if size != 0 {
			t.Fatalf("after size is %d", size)
		}
	})
}

func TestDelete2(t *testing.T) {
	tree := NewTree[int, string]()
	tree.Insert(4, "1qa")
	tree.Insert(2, "2ws")
	tree.Insert(3, "3ed")
	tree.Insert(1, "4rf")
	tree.Insert(8, "5tg")
	tree.Insert(5, "6yh")
	tree.Insert(7, "7uj")
	tree.Insert(9, "8ik")
	tree.Delete(1)
	tree.Delete(2)
}
