package oam

type iterator[K comparable, V any] struct {
	m *M[K, V]
	i int
}

func newIterator[K comparable, V any](m *M[K, V]) *iterator[K, V] {
	return &iterator[K, V]{
		m: m,
		i: -1,
	}
}

func (iter *iterator[K, V]) Next() (K, bool) {
	k, ok, i := iter.m.next(iter.i + 1)
	if ok {
		iter.i = i
	}

	return k, ok
}

func (iter *iterator[K, V]) Prev() (K, bool) {
	k, ok, i := iter.m.prev(iter.i - 1)
	if ok {
		iter.i = i
	}

	return k, ok
}

func (iter *iterator[K, V]) Begin() {
	iter.i = iter.m.begin()
}

func (iter *iterator[K, V]) End() {
	iter.i = iter.m.end()
}
