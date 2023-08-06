package oam

import "sync"

type M[K comparable, V any] struct {
	mx          sync.RWMutex
	data        map[K]V
	index       map[K]int
	alive       []bool
	keys        []K
	concurrency bool
}

func New[K comparable, V any](capcity int, concurrency bool) *M[K, V] {
	return &M[K, V]{
		data:        make(map[K]V, capcity),
		index:       make(map[K]int, capcity),
		alive:       make([]bool, 0, capcity),
		keys:        make([]K, 0, capcity),
		concurrency: concurrency,
	}
}

func (m *M[K, V]) Add(k K, v V) {
	if m.concurrency {
		m.mx.Lock()
		defer m.mx.Unlock()
	}

	m.data[k] = v
	m.alive = append(m.alive, true)
	m.keys = append(m.keys, k)
	m.index[k] = len(m.alive) - 1
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
