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

type node struct {
	str string
	vl  int
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

	fmt.Println("Find test begin")
	n := tree.FindIt(key(4))
	fmt.Println(n.Value)
	n.Value = "bdsf"
	fmt.Println(n.Value)
	n = tree.FindIt(key(4))
	fmt.Println(n.Value)
	fmt.Println(tree.Find(key(5)))
	fmt.Println("Find test end")

	fmt.Println("Iterator test begin")

	it := tree.Iterator()

	for it != nil {
		fmt.Println(it.Value)
		fmt.Println(it.Key)
		it = it.Next()
	}

	fmt.Println("Iterator test end")

	fmt.Println("Type test begin")
	v := tree.Find(key(4))
	fmt.Printf("%T\n", v)
	tree2 := rbtree.NewTree()
	tree2.Insert(key(1), node{"ads", 1})

	no, _ := tree2.Find(key(1)).(node)
	fmt.Printf("%T\n", no)
	fmt.Println("Type test end")

	fmt.Println("Delete test begin")
	tree.Delete(key(6))
	tree.Preorder()
	for i := 1; i <= 6; i++ {
		fmt.Println(i)
		tree.Delete(key(i))
		tree.Preorder()
		fmt.Printf("size is %d\n", tree.Size())
	}
	fmt.Println("Delete test end")

	fmt.Println("Clear test begin")
	tree2.Clear()
	fmt.Println(tree2.Find(key(1)))
	fmt.Println("Clear test end")
}
