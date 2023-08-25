package om

import (
	"sync"

	"github.com/nastvood/om/conf"
	"golang.org/x/exp/constraints"
)

type M[K constraints.Ordered, V any] struct {
	mx          sync.RWMutex
	data        map[K]V
	tree        *tree[K]
	concurrency bool
}

func New[K constraints.Ordered, V any](opts ...conf.Option) *M[K, V] {
	c := conf.DefaultConfig()
	for _, opt := range opts {
		opt(&c)
	}

	return &M[K, V]{
		data:        make(map[K]V, c.Capacity),
		tree:        newTree[K](c.BucketLen),
		concurrency: c.Concurrency,
	}
}

func (m *M[K, V]) Add(k K, v V) {
	if m.concurrency {
		m.mx.Lock()
		defer m.mx.Unlock()
	}

	_, ok := m.data[k]
	if !ok {
		m.tree.insert(k)
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
		m.tree.delete(k)
	}
}

func (m *M[K, V]) Keys() []K {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	return m.tree.inorder()
}
