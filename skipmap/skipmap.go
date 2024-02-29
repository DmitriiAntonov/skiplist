package skipmap

import (
	"github.com/dmitriiantonov/skiplist/comparator"
	"math/rand"
	"time"
)

const (
	maxLevel         = 16
	p        float64 = 0.5
)

type node[K any, V any] struct {
	key  K
	val  V
	next []*node[K, V]
}

func newNode[K any, V any](key K, val V, level int) *node[K, V] {
	return &node[K, V]{
		key:  key,
		val:  val,
		next: make([]*node[K, V], level),
	}
}

type SkipMap[K any, V any] struct {
	head       *node[K, V]
	comparator comparator.Comparator[K]
	level      int
	rand       *rand.Rand
}

// New creates a new instance of SkipMap
func New[K any, V any](comparator comparator.Comparator[K]) *SkipMap[K, V] {
	return &SkipMap[K, V]{
		head:       newNode(*new(K), *new(V), maxLevel),
		comparator: comparator,
		level:      1,
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m *SkipMap[K, V]) randomLevel() int {
	level := 1
	for m.rand.Float64() < p && level < maxLevel {
		level++
	}
	return level
}

// Insert key, value into skip map
func (m *SkipMap[K, V]) Insert(key K, val V) {
	update := make([]*node[K, V], maxLevel)
	current := m.head

	for i := m.level - 1; i >= 0; i-- {
		for current.next[i] != nil && m.comparator.Compare(current.next[i].key, key) < 0 {
			current = current.next[i]
		}
		update[i] = current
	}

	current = current.next[0]

	if current != nil && m.comparator.Compare(current.key, key) == 0 {
		current.val = val
		return
	}

	if current == nil || m.comparator.Compare(current.key, key) != 0 {
		newLevel := m.randomLevel()

		if newLevel > m.level {
			for i := m.level; i < newLevel; i++ {
				update[i] = m.head
			}
			m.level = newLevel
		}

		insertedNode := newNode(key, val, newLevel)

		for i := 0; i < newLevel; i++ {
			insertedNode.next[i] = update[i].next[i]
			update[i].next[i] = insertedNode
		}
	}
}

// Get a value by the key
func (m *SkipMap[K, V]) Get(key K) (V, bool) {
	current := m.head

	for i := m.level - 1; i >= 0; i-- {
		for current.next[i] != nil && m.comparator.Compare(current.next[i].key, key) < 0 {
			current = current.next[i]
		}
	}

	current = current.next[0]

	if current != nil && m.comparator.Compare(current.key, key) == 0 {
		return current.val, true
	}

	return *new(V), false
}
