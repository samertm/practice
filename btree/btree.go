package btree

import (
	"strconv"
)

type Btree struct {
	Item   int
	Parent *Btree
	Left   *Btree
	Right  *Btree
}

func New(val int) *Btree {
	return &Btree{Item: val}
}

func (t *Btree) Search(item int) *Btree {
	if t.Item == item {
		return t
	}
	if item < t.Item {
		if t.Left == nil {
			return nil
		}
		return Search(t.Left, item)
	}
	if t.Right == nil {
		return nil
	}
	return Search(t.Right, item)
}

func (t *Btree) FindMinimum() *Btree {
	for min := t; min.Left != nil; min = min.Left {
		// this block left intentionally blank
	}
	return min
}

func (t *Btree) Insert(i int) {
	if t.Item == i {
		return
	}
	if i < t.Item {
		if t.Left != nil {
			t.Left.Insert(i)
			return
		}
		t.Left = &Btree{Item: i, Parent: t}
	} else {
		if t.Right != nil {
			t.Right.Insert(i)
		}
		t.Right = &Btree{Item: i, Parent: t}
	}
}

func (t *Btree) String() (s string) {
	if t.Left != nil {
		s += t.Left.String()
	}
	s += strconv.Itoa(t.Item)
	if t.Right != nil {
		s += t.Right.String()
	}
	return
}
