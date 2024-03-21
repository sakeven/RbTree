package rbtree

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	"testing"
)

func BenchmarkRBTree(b *testing.B) {
	b.Run("insert", func(b *testing.B) {
		tree := NewTree[int64, string]()
		keys, values := createData(b.N)

		b.ReportAllocs()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			tree.Insert(keys[i], values[i])
		}
	})

	b.Run("iterate", func(b *testing.B) {
		tree := NewTree[int64, string]()
		keys, values := createData(b.N)

		for i := 0; i < b.N; i++ {
			tree.Insert(keys[i], values[i])
		}

		length := tree.Size()

		b.ReportAllocs()
		b.ResetTimer()

		it := tree.Iterator()
		for i := 0; i < length-1; i++ {
			it = it.Next()
			if it == nil {
				b.Fatalf("%d call to .Iterator() returned nil", i)
			}
		}
	})

	b.Run("delete", func(b *testing.B) {
		tree := NewTree[int64, string]()
		keys, values := createData(b.N)

		for i := 0; i < b.N; i++ {
			tree.Insert(keys[i], values[i])
		}

		length := int64(tree.Size())

		b.ReportAllocs()
		b.ResetTimer()

		var i int64
		for i = 0; i < length-1; i++ {
			tree.Delete(i)
		}
	})
}

func createData(iterations int) ([]int64, []string) {
	max := big.NewInt(math.MaxInt)

	keys := make([]int64, iterations)
	for i := 0; i < iterations; i++ {
		n, _ := rand.Int(rand.Reader, max)
		keys[i] = n.Int64()
	}
	values := make([]string, iterations)
	for i := 0; i < iterations; i++ {
		values[i] = fmt.Sprintf("%20.20d", keys[i])
	}

	return keys, values
}
