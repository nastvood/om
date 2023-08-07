package om

import (
	"sync"

	"golang.org/x/exp/constraints"
)

type M[K constraints.Ordered, V any] struct {
	mx          sync.RWMutex
	data        map[K]V
	tree        *node[K]
	concurrency bool
}

func New[K constraints.Ordered, V any](capacity int, concurrency bool) *M[K, V] {
	return &M[K, V]{
		data:        make(map[K]V, capacity),
		concurrency: concurrency,
	}
}

func (m *M[K, V]) Add(k K, v V) {
	if m.concurrency {
		m.mx.Lock()
		defer m.mx.Unlock()
	}

	_, ok := m.data[k]
	if !ok {
		m.tree = insert(m.tree, k)
	}

	m.data[k] = v
}

func (m *M[K, V]) Get(k K) (V, bool) {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	v, ok := m.data[k]

	return v, ok
}

func (m *M[K, V]) Delete(k K) {
	if m.concurrency {
		m.mx.Lock()
		defer m.mx.Unlock()
	}

	_, ok := m.data[k]
	if ok {
		m.tree = deleteNode(m.tree, k)
	}
}
