package main

import (
	"github.com/samertm/practice/btree"
	"fmt"
)

func main() {
	t := btree.New(6)
	t.Insert(3)
	t.Insert(55)
	t.Insert(56)
	fmt.Println(t)
}
