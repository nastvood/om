package om

import (
	"fmt"

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
	n := find(root, k)
	if n == nil {
		return root
	}

	// нет потомков
	if n.right == nil && n.left == nil {
		// если корень
		if n.parent == nil {
			return nil
		}

		if n == n.parent.left {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}

		return root
	}

	var x, y *node[K]

	deletedIsRed := n.isRed

	// один потомок
	// то делаем у родителя удаляемой вершины ссылку на этого потомка вместо удаляемой вершины
	if n.left == nil || n.right == nil {
		if n.left == nil {
			x = n.right
			n.right = nil
		} else {
			x = n.left
			n.left = nil
		}

		if n == n.parent.left {
			n.parent.left = x
		} else {
			n.parent.right = x
		}

		x.parent = n.parent

		n.parent = nil
	} else {
		// два потомка

		// у - не имеет левого потомка, у - ближайшее (большее) значение к n
		y = n.right
		for y.left != nil {
			y = y.left
		}
		deletedIsRed = y.isRed

		x = y.right
		if y.parent == n {
			// предок у это удаляемая вершина, тогда присоединяем правого потомка y к удаляемой вершине
			x.parent = n
		} else {
			// заменяем y его правым потомком
			y.parent.left = x
			x.parent = y.parent
		}

		y.parent = nil
		y.right = nil

		n.data = y.data
	}

	if !deletedIsRed {
		return deleteFixup(root, x)
	}

	return root
}

func deleteFixup[K constraints.Ordered](root *node[K], n *node[K]) *node[K] {
	fmt.Println(n)
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
