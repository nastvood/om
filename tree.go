package om

import "golang.org/x/exp/constraints"

type tree[K constraints.Ordered] struct {
	root  *node[K]
	nodes []node[K]
}

func newTree[K constraints.Ordered](capacity int) *tree[K] {
	return &tree[K]{
		nodes: make([]node[K], 0, capacity),
	}
}

func (t *tree[K]) insert(k K) {
	t.nodes = append(t.nodes, node[K]{
		data:  k,
		isRed: true,
	})

	t.root = insert(t.root, &t.nodes[len(t.nodes)-1])
}

/*func (t *tree[K]) exists(k K) bool {
	if n := find(t.root, k); n != nil && !n.isDel {
		return true
	}

	return false
}*/

func (t *tree[K]) delete(k K) {
	if n := find(t.root, k); n != nil {
		n.isDel = true
	}
}

func (t *tree[K]) inorder() []K {
	return t.root.inorder()
}
