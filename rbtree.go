package rbtree

// color of node
const (
	RED   = 0
	BLACK = 1
)

type Keytype interface {
	LessThan(interface{}) bool
}

type valuetype interface{}

type node struct {
	left, right, parent *node
	color               int
	Key                 Keytype
	Value               valuetype
}

// Tree is a struct of red-black tree.
type Tree struct {
	root *node
	size int
}

// NewTree creates a new rbtree.
func NewTree() *Tree {
	return &Tree{}
}

// Find finds the node and return its value.
func (t *Tree) Find(key Keytype) interface{} {
	n := t.findnode(key)
	if n != nil {
		return n.Value
	}
	return nil
}

// FindIt finds the node and return it as an iterator.
func (t *Tree) FindIt(key Keytype) *node {
	return t.findnode(key)
}

// Empty checks whether the rbtree is empty.
func (t *Tree) Empty() bool {
	if t.root == nil {
		return true
	}
	return false
}

// Iterator creates the rbtree's iterator that points to the minmum node.
func (t *Tree) Iterator() *node {
	return minimum(t.root)
}

// Size returns the size of the rbtree.
func (t *Tree) Size() int {
	return t.size
}

// Clear destroys the rbtree.
func (t *Tree) Clear() {
	t.root = nil
	t.size = 0
}

// Insert inserts the key-value pair into the rbtree.
func (t *Tree) Insert(key Keytype, value valuetype) {
	x := t.root
	var y *node

	for x != nil {
		y = x
		if key.LessThan(x.Key) {
			x = x.left
		} else {
			x = x.right
		}
	}

	z := &node{parent: y, color: RED, Key: key, Value: value}
	t.size++

	if y == nil {
		z.color = BLACK
		t.root = z
		return
	} else if z.Key.LessThan(y.Key) {
		y.left = z
	} else {
		y.right = z
	}
	t.rbInsertFixup(z)

}

// Delete deletes the node by key
func (t *Tree) Delete(key Keytype) {
	z := t.findnode(key)
	if z == nil {
		return
	}

	var x, y *node
	if z.left != nil && z.right != nil {
		y = successor(z)
	} else {
		y = z
	}

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	xparent := y.parent
	if x != nil {
		x.parent = xparent
	}
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.left {
		y.parent.left = x
	} else {
		y.parent.right = x
	}

	if y != z {
		z.Key = y.Key
		z.Value = y.Value
	}

	if y.color == BLACK {
		t.rbDeleteFixup(x, xparent)
	}
	t.size--
}

func (t *Tree) rbInsertFixup(z *node) {
	var y *node
	for z.parent != nil && z.parent.color == RED {
		if z.parent == z.parent.parent.left {
			y = z.parent.parent.right
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.rightRotate(z.parent.parent)
			}
		} else {
			y = z.parent.parent.left
			if y != nil && y.color == RED {
				z.parent.color = BLACK
				y.color = BLACK
				z.parent.parent.color = RED
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = BLACK
				z.parent.parent.color = RED
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

func (t *Tree) rbDeleteFixup(x, parent *node) {
	var w *node

	for x != t.root && getColor(x) == BLACK {
		if x != nil {
			parent = x.parent
		}
		if x == parent.left {
			w = parent.right
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.leftRotate(parent)
				w = parent.right
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.right) == BLACK {
					if w.left != nil {
						w.left.color = BLACK
					}
					w.color = RED
					t.rightRotate(w)
					w = parent.right
				}
				w.color = parent.color
				parent.color = BLACK
				if w.right != nil {
					w.right.color = BLACK
				}
				t.leftRotate(parent)
				x = t.root
			}
		} else {
			w = parent.left
			if w.color == RED {
				w.color = BLACK
				parent.color = RED
				t.rightRotate(parent)
				w = parent.left
			}
			if getColor(w.left) == BLACK && getColor(w.right) == BLACK {
				w.color = RED
				x = parent
			} else {
				if getColor(w.left) == BLACK {
					if w.right != nil {
						w.right.color = BLACK
					}
					w.color = RED
					t.leftRotate(w)
					w = parent.left
				}
				w.color = parent.color
				parent.color = BLACK
				if w.left != nil {
					w.left.color = BLACK
				}
				t.rightRotate(parent)
				x = t.root
			}
		}
	}
	if x != nil {
		x.color = BLACK
	}
}

func (t *Tree) leftRotate(x *node) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func (t *Tree) rightRotate(x *node) {
	y := x.left
	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.right = x
	x.parent = y
}

// findnode finds the node by key and return it, if not exists return nil.
func (t *Tree) findnode(key Keytype) *node {
	x := t.root
	for x != nil {
		if key.LessThan(x.Key) {
			x = x.left
		} else {
			if key == x.Key {
				return x
			}
			x = x.right
		}
	}
	return nil
}

// Next returns the node's successor as an iterator.
func (n *node) Next() *node {
	return successor(n)
}

// successor returns the successor of the node
func successor(x *node) *node {
	if x.right != nil {
		return minimum(x.right)
	}
	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = x.parent
	}
	return y
}

// getColor gets color of the node.
func getColor(n *node) int {
	if n == nil {
		return BLACK
	}
	return n.color
}

// minimum finds the minimum node of subtree n.
func minimum(n *node) *node {
	for n.left != nil {
		n = n.left
	}
	return n
}
