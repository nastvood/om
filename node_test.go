package om

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func Test_insert_down_to(t *testing.T) {
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
}

func Test_insert_up(t *testing.T) {
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
}

func Test_insert_rand(t *testing.T) {
	for _, cnt := range []int{1, 10, 50, 100, 150, 200} {
		vals, tree := testIntRandSliceToTree(t, cnt)
		slices.Sort(vals)
		require.ElementsMatch(t, vals, tree.inorder())
	}
}

func testIntRandSliceToTree(t *testing.T, n int) ([]int, *node[int]) {
	t.Helper()

	cnt := 10

	var tree *node[int]

	vals := rand.Perm(cnt)
	for _, v := range vals {
		tree = insert(tree, v)
	}

	return vals, tree
}
