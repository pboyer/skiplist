package main

import (
	"fmt"
	"math/rand"

	"github.com/pboyer/skiplist"
)

const MAX_HEIGHT = 6

func main() {
	list := skiplist.New(MAX_HEIGHT)

	for i := 0; i < 50; i++ {
		n := rand.Intn(10000)
		fmt.Println("inserting", n)
		list.Put(n, struct{}{})
	}

	list.Dump()

	node, ok := list.Get(8070)
	fmt.Println(node, ok)
	fmt.Println(list.Remove(node))
	node, ok = list.Get(8070)
	fmt.Println(node, ok)
}
