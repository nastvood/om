package om

import (
	"golang.org/x/exp/constraints"
)

type node[K constraints.Ordered] struct {
	parent, left, right *node[K]
	isRed               bool
	data                K
}

func insert[K constraints.Ordered](root *node[K], k K) *node[K] {
	current := root
	var parent *node[K]

	for current != nil {
		parent = current
		if k < current.data {
			current = current.left
		} else {
			current = current.right
		}
	}

	n := &node[K]{
		data:   k,
		parent: parent,
		isRed:  true,
	}

	if parent == nil {
		return n
	}

	if k < parent.data {
		parent.left = n
	} else {
		parent.right = n
	}

	return insertFixup[K](root, n)
}

func insertFixup[K constraints.Ordered](root *node[K], n *node[K]) *node[K] {
	// вставили красный и родитель красный
	for n != root && n.parent.isRed && n.parent.parent != nil {
		// родитель слева
		if n.parent == n.parent.parent.left {
			r := n.parent.parent.right
			if r != nil && r.isRed {
				// 1. дядя красный
				n.parent.isRed = false
				r.isRed = false
				n.parent.parent.isRed = false
				n = n.parent.parent
			} else {
				// 2. дядя черный
				// если добавляемый узел был правым потомком сначала делаем левое вращение
				if n == n.parent.right {
					n = n.parent
					root = rotateLeft(root, n)
				}

				n.parent.isRed = false
				n.parent.parent.isRed = true
				root = rotateRight(root, n.parent.parent)
			}
		} else {
			// родитель справа
			l := n.parent.parent.left
			if l != nil && l.isRed {
				// 3. дядя красный
				n.parent.isRed = false
				l.isRed = false
				n.parent.parent.isRed = true
				n = n.parent.parent
			} else {
				// 4. дядя черный
				// если добавляемый узел был левым потомком сначала делаем правое вращение
				if n == n.parent.left {
					n = n.parent
					root = rotateRight(root, n)
				}

				n.parent.isRed = false
				n.parent.parent.isRed = true
				root = rotateLeft(root, n.parent.parent)
			}
		}
	}

	root.isRed = false

	return root
}

func find[K constraints.Ordered](root *node[K], k K) *node[K] {
	current := root
	for current != nil {
		if current.data == k {
			return current
		}

		if current.data > k {
			current = current.left
			continue
		}

		current = current.right
	}

	return nil
}

func deleteNode[K constraints.Ordered](root *node[K], k K) *node[K] {
	z := find(root, k)
	if z == nil {
		return root
	}

	// нет потомков
	if z.right == nil && z.left == nil {
		// если корень
		if z.parent == nil {
			return nil
		}

		if z == z.parent.left {
			z.parent.left = nil
		} else {
			z.parent.right = nil
		}

		z.parent = nil

		return root
	}

	var x *node[K]

	deletedIsRed := z.isRed

	// один потомок
	// то делаем у родителя удаляемой вершины ссылку на этого потомка вместо удаляемой вершины
	if z.left == nil || z.right == nil {
		if z.left == nil {
			x = z.right
			z.right = nil
		} else {
			x = z.left
			z.left = nil
		}

		if z == z.parent.left {
			z.parent.left = x
		} else {
			z.parent.right = x
		}

		x.parent = z.parent

		z.parent = nil
	} else {
		// два потомка

		// Y - не имеет левого потомка, Y - ближайшее (большее) значение к z
		y := z.right
		for y.left != nil {
			y = y.left
		}
		deletedIsRed = y.isRed

		x = y.right
		if y.parent == z && x != nil {
			// предок Y это удаляемая вершина, тогда присоединяем правого потомка Y к удаляемой вершине
			x.parent = z
		} else {
			// заменяем Y его правым потомком
			if y.parent != nil {
				y.parent.left = x
			}
			if x != nil {
				x.parent = y.parent
			}
		}

		y.parent = nil
		y.right = nil

		z.data = y.data
	}

	if !deletedIsRed {
		return deleteFixup(root, x)
	}

	return root
}

func deleteFixup[K constraints.Ordered](root *node[K], n *node[K]) *node[K] {
	if n == nil {
		return root
	}

	for n != root && !n.isRed {
		if n == n.parent.left {
			w := n.parent.right
			if w.isRed {
				w.isRed = false
				n.parent.isRed = true
				root = rotateLeft(root, n.parent)
				w = n.parent.right
			}

			if !w.left.isRed && !w.right.isRed {
				w.isRed = true
				n = n.parent
			} else {
				if !w.right.isRed {
					w.left.isRed = false
					w.isRed = true
					root = rotateRight(root, w)
					w = n.parent.right
				}

				w.isRed = n.parent.isRed
				n.parent.isRed = false
				w.right.isRed = false
				n = rotateLeft(root, n.parent)
			}
		} else {
			w := n.parent.left
			if w.isRed {
				w.isRed = false
				n.parent.isRed = true
				root = rotateRight(root, n.parent)
				w = n.parent.left
			}

			if !w.left.isRed && !w.right.isRed {
				w.isRed = true
				n = n.parent
			} else {
				if !w.left.isRed {
					w.right.isRed = false
					w.isRed = true
					root = rotateLeft(root, w)
					w = n.parent.left
				}

				w.isRed = n.parent.isRed
				n.parent.isRed = false
				w.left.isRed = false
				n = rotateRight(root, n.parent)
			}
		}
	}

	n.isRed = false

	return root
}

func rotateLeft[K constraints.Ordered](root *node[K], x *node[K]) *node[K] {
	y := x.right

	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	// x - root
	if x.parent == nil {
		root = y
	} else {
		if x == x.parent.left {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	y.left = x
	x.parent = y

	return root
}

func rotateRight[K constraints.Ordered](root *node[K], x *node[K]) *node[K] {
	y := x.left

	x.left = y.right
	if y.right != nil {
		y.right.parent = x
	}

	y.parent = x.parent

	// x - root
	if x.parent == nil {
		root = y
	} else {
		if x == x.parent.right {
			x.parent.right = y
		} else {
			x.parent.left = y
		}
	}

	y.right = x
	x.parent = y

	return root
}

func fetch[K constraints.Ordered](data []K, x *node[K]) []K {
	if x == nil {
		return data
	}

	data = fetch(data, x.left)
	data = append(data, x.data)
	data = fetch(data, x.right)

	return data
}

func (x *node[K]) inorder() []K {
	if x == nil {
		return nil
	}

	return fetch(make([]K, 0), x)
}
