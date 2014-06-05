package main

import (
	"fmt"
	"rbtree"
)

type key int

func (n key) LessThan(b interface{}) bool {
	value, _ := b.(key)
	return n < value
}

func main() {
	tree := rbtree.NewTree()

	tree.Insert(key(1), "123")
	tree.Insert(key(3), "234")
	tree.Insert(key(4), "dfa3")
	tree.Insert(key(6), "sd4")
	tree.Insert(key(5), "jcd4")
	tree.Insert(key(2), "bcd4")
	fmt.Printf("size is %d\n", tree.Size())
	tree.Preorder()

	fmt.Println(tree.Find(key(5)))
	tree.Delete(key(6))
	tree.Preorder()
	for i := 1; i <= 6; i++ {
		fmt.Println(i)
		tree.Delete(key(i))
		tree.Preorder()
		fmt.Printf("size is %d\n", tree.Size())
	}
}
