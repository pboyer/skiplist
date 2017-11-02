package skiplist

import (
	"math/rand"
	"testing"
)

const HEIGHT = 16

func benchPut(size int, b *testing.B) {
	list := New(HEIGHT)
	for n := 0; n < b.N; n++ {
		for i := 0; i < size; i++ {
			list.Put(n, struct{}{})
		}
	}
}

func BenchmarkPut100(b *testing.B) {
	benchPut(100, b)
}

func BenchmarkPut1000(b *testing.B) {
	benchPut(1000, b)
}

func BenchmarkPut10000(b *testing.B) {
	benchPut(10000, b)
}

func benchPutRemoveAll(size int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		ns := []*Node{}
		list := New(HEIGHT)
		for i := 0; i < size; i++ {
			ns = append(ns, list.Put(i, struct{}{}))
		}

		for _, n := range ns {
			list.Remove(n)
		}
	}
}

func BenchmarkPutRemoveAll100(b *testing.B) {
	benchPutRemoveAll(100, b)
}

func BenchmarkPutRemoveAll1000(b *testing.B) {
	benchPutRemoveAll(1000, b)
}

func BenchmarkPutRemoveAll10000(b *testing.B) {
	benchPutRemoveAll(10000, b)
}

func BenchmarkPutRemove(b *testing.B) {
	b.StopTimer()
	list := New(HEIGHT)

	for i := 0; i < 1000; i++ {
		list.Put(rand.Intn(10000), struct{}{})
	}
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		node := list.Put(n, struct{}{})
		list.Remove(node)
	}
}
func TestPutRemoveAll(t *testing.T) {
	size := 100

	ns := []*Node{}
	list := New(HEIGHT)
	for i := 0; i < size; i++ {
		ns = append(ns, list.Put(rand.Intn(size*size), struct{}{}))
	}

	for _, n := range ns {
		_, ok := list.Get(n.Key)
		if !ok {
			t.Fatalf("Failed to find node just inserted: %d", n)
		}
	}

	for _, n := range ns {
		ok := list.Remove(n)
		if !ok {
			t.Fatalf("Failed to remove node just inserted: %d", n)
		}
	}
}
