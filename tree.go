package om

import (
	"golang.org/x/exp/constraints"
)

type tree[K constraints.Ordered] struct {
	root      *node[K]
	buckets   [][]node[K]
	deleted   map[K]bucketPos
	curBucket int
}

const bucketLen = 100

func newTree[K constraints.Ordered]() *tree[K] {
	return &tree[K]{
		buckets: [][]node[K]{make([]node[K], 0, bucketLen)},
		deleted: make(map[K]bucketPos),
	}
}

func (t *tree[K]) insert(k K) {
	var n *node[K]

	pos, ok := t.deleted[k]
	if ok {
		n = &t.buckets[pos.row][pos.col]
		n.isDel = false
		t.root = insert(t.root, n)
		delete(t.deleted, k)

		return
	}

	if len(t.buckets[t.curBucket]) == bucketLen {
		t.buckets = append(t.buckets, make([]node[K], 0, bucketLen))
		t.curBucket++
	}

	t.buckets[t.curBucket] = append(t.buckets[t.curBucket], node[K]{
		data: k,
		pos: bucketPos{
			row: t.curBucket,
			col: len(t.buckets[t.curBucket]),
		},
		isRed: true,
	})

	t.root = insert(t.root, &t.buckets[t.curBucket][len(t.buckets[t.curBucket])-1])
}

func (t *tree[K]) delete(k K) {
	if n := find(t.root, k); n != nil && !n.isDel {
		t.deleted[k] = n.pos
		n.isDel = true
	}
}

func (t *tree[K]) inorder() []K {
	return t.root.inorder()
}
