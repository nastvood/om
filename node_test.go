package om

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func Test_insert(t *testing.T) {
	t.Run("down to", func(t *testing.T) {
		var tree *node[int]

		tree = insert(tree, 9)
		require.ElementsMatch(t, []int{9}, tree.inorder())
		tree = insert(tree, 8)
		require.ElementsMatch(t, []int{8, 9}, tree.inorder())
		tree = insert(tree, 7)
		require.ElementsMatch(t, []int{7, 8, 9}, tree.inorder())
		tree = insert(tree, 6)
		require.ElementsMatch(t, []int{6, 7, 8, 9}, tree.inorder())
		tree = insert(tree, 5)
		require.ElementsMatch(t, []int{5, 6, 7, 8, 9}, tree.inorder())
		tree = insert(tree, 4)
		require.ElementsMatch(t, []int{4, 5, 6, 7, 8, 9}, tree.inorder())
		tree = insert(tree, 3)
		require.ElementsMatch(t, []int{3, 4, 5, 6, 7, 8, 9}, tree.inorder())
		tree = insert(tree, 2)
		require.ElementsMatch(t, []int{2, 3, 4, 5, 6, 7, 8, 9}, tree.inorder())
		tree = insert(tree, 1)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tree.inorder())
	})

	t.Run("up", func(t *testing.T) {
		var tree *node[int]

		tree = insert(tree, 1)
		require.ElementsMatch(t, []int{1}, tree.inorder())
		tree = insert(tree, 2)
		require.ElementsMatch(t, []int{1, 2}, tree.inorder())
		tree = insert(tree, 3)
		require.ElementsMatch(t, []int{1, 2, 3}, tree.inorder())
		tree = insert(tree, 4)
		require.ElementsMatch(t, []int{1, 2, 3, 4}, tree.inorder())
		tree = insert(tree, 5)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5}, tree.inorder())
		tree = insert(tree, 6)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, tree.inorder())
		tree = insert(tree, 7)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7}, tree.inorder())
		tree = insert(tree, 8)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, tree.inorder())
		tree = insert(tree, 9)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tree.inorder())
	})

	t.Run("rand", func(t *testing.T) {
		for _, cnt := range []int{0, 1, 10, 50, 100, 150, 200} {
			t.Run(strconv.Itoa(cnt), func(t *testing.T) {
				vals, tree := testIntRandSliceToTree(t, cnt)
				slices.Sort(vals)
				require.ElementsMatch(t, vals, tree.inorder())
			})
		}
	})
}

func Test_find(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var empty *node[int]
		v := find(empty, 15)
		require.Nil(t, v)
	})

	t.Run("rand", func(t *testing.T) {
		cnt := 10
		vals, tree := testIntRandSliceToTree(t, cnt)

		v := find(tree, cnt+1)
		require.Nil(t, v)

		v = find(tree, 3)
		require.NotNilf(t, v, "vals %v, key %d", vals, 3)

		for _, k := range vals {
			n := find(tree, k)
			require.NotNilf(t, n, "vals %v, key %d", vals, k)
			require.Equal(t, k, n.data)
		}
	})
}

func Test_delete(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var empty *node[int]
		empty = deleteNode(empty, 10)
		require.Nil(t, empty)
	})

	t.Run("one", func(t *testing.T) {
		var tree *node[int]
		tree = insert(tree, 1)
		tree = deleteNode(tree, 10)
		require.NotNil(t, tree)
		tree = deleteNode(tree, 1)
		require.Nil(t, tree)
	})

	t.Run("three", func(t *testing.T) {
		var tree *node[int]
		tree = insert(tree, 2)
		tree = insert(tree, 1)
		tree = insert(tree, 3)

		tree = deleteNode(tree, 10)
		require.NotNil(t, tree)

		tree = deleteNode(tree, 1)
		require.ElementsMatch(t, []int{2, 3}, tree.inorder())

		tree = deleteNode(tree, 3)
		require.ElementsMatch(t, []int{2}, tree.inorder())

		tree = deleteNode(tree, 2)
		require.Nil(t, tree)

	})
}

func testIntRandSliceToTree(t *testing.T, n int) ([]int, *node[int]) {
	t.Helper()

	var tree *node[int]

	vals := rand.Perm(n)
	for _, v := range vals {
		tree = insert(tree, v)
	}

	return vals, tree
}
