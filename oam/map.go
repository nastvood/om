package oam

import (
	"sync"

	"github.com/nastvood/om/conf"
)

type M[K comparable, V any] struct {
	mx          sync.RWMutex
	data        map[K]V
	index       map[K]int
	alive       []bool
	keys        []K
	concurrency bool
}

func New[K comparable, V any](opts ...conf.Option) *M[K, V] {
	c := conf.DefaultConfig()
	for _, opt := range opts {
		opt(&c)
	}

	return &M[K, V]{
		data:        make(map[K]V, c.Capacity),
		index:       make(map[K]int, c.Capacity),
		alive:       make([]bool, 0, c.Capacity),
		keys:        make([]K, 0, c.Capacity),
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
		m.alive = append(m.alive, true)
		m.keys = append(m.keys, k)
		m.index[k] = len(m.alive) - 1
	}

	m.data[k] = v
}

func (m *M[K, V]) Get(k K) (V, bool) {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	_, ok := m.index[k]
	return m.data[k], ok
}

func (m *M[K, V]) Delete(k K) {
	if m.concurrency {
		m.mx.Lock()
		defer m.mx.Unlock()
	}

	i, ok := m.index[k]
	if !ok {
		return
	}

	delete(m.index, k)
	delete(m.data, k)
	m.alive[i] = false
}

func (m *M[K, V]) Len() int {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	return len(m.data)
}

//nolint:revive
func (m *M[K, V]) Iterator() *iterator[K, V] {
	return newIterator[K, V](m)
}

func (m *M[K, V]) next(i int) (K, bool, int) {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	for {
		if i < 0 || i >= len(m.alive) {
			var k K
			return k, false, i
		}

		if !m.alive[i] {
			i++
			continue
		}

		return m.keys[i], true, i
	}
}

func (m *M[K, V]) prev(i int) (K, bool, int) {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	for {
		if i < 0 || i >= len(m.alive) {
			var k K
			return k, false, i
		}

		if !m.alive[i] {
			i--
			continue
		}

		return m.keys[i], true, i
	}
}

func (m *M[K, V]) begin() int {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	i := 0

	for {
		if i >= len(m.alive) {
			return -1
		}

		if !m.alive[i] {
			i++
			continue
		}

		return i - 1
	}
}

func (m *M[K, V]) end() int {
	if m.concurrency {
		m.mx.RLock()
		defer m.mx.RUnlock()
	}

	i := len(m.alive) - 1

	for {
		if i <= 0 {
			return 0
		}

		if !m.alive[i] {
			i--
			continue
		}

		return i + 1
	}
}
