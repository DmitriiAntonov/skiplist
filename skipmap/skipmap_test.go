package skipmap

import (
	"github.com/dmitriiantonov/skiplist/comparator"
	"testing"
)

func TestInsert(t *testing.T) {

	list := New[int, int](comparator.NewDefaultComparator[int]())

	pairs := make([]int, 1000)

	for i := range pairs {
		pairs[i] = i * 2
	}

	for key, value := range pairs {
		if key%2 == 0 || key%3 == 0 {
			list.Insert(key, value)
		}
	}

	for key, want := range pairs {
		if key%2 == 0 || key%3 == 0 {
			if got, _ := list.Get(key); got != want {
				t.Errorf("Get(%d): got %d, want %d", key, got, want)
			}
		} else {
			if _, ok := list.Get(key); ok {
				t.Errorf("Get(%d): key found in list", key)
			}
		}
	}
}

func TestUpdate(t *testing.T) {

	list := New[int, int](comparator.NewDefaultComparator[int]())

	pairs := make([]int, 100)

	for i := range pairs {
		pairs[i] = i * 2
	}

	for key, value := range pairs {
		list.Insert(key, value)
	}

	for key, want := range pairs {
		if got, _ := list.Get(key); got != want {
			t.Errorf("Get(%d): got %d, want %d", key, got, want)
		}
	}

	for i := range pairs {
		pairs[i] = i * 4
	}

	for key, value := range pairs {
		list.Insert(key, value)
	}

	for key, want := range pairs {
		if got, _ := list.Get(key); got != want {
			t.Errorf("Get(%d): got %d, want %d", key, got, want)
		}
	}
}
