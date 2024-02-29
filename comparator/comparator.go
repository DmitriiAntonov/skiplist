package comparator

import "cmp"

type Comparator[K any] interface {
	Compare(a, b K) int
}

type DefaultComparator[K cmp.Ordered] struct {
}

func NewDefaultComparator[K cmp.Ordered]() *DefaultComparator[K] {
	return new(DefaultComparator[K])
}

func (c *DefaultComparator[K]) Compare(x, y K) int {
	return cmp.Compare(x, y)
}
