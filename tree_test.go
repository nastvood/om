package om

import (
	"strconv"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func Test_insert(t *testing.T) {
	t.Run("down to", func(t *testing.T) {
		tr := newTree[int](0)

		tr.insert(9)
		require.ElementsMatch(t, []int{9}, tr.inorder())
		tr.insert(8)
		require.ElementsMatch(t, []int{8, 9}, tr.inorder())
		tr.insert(7)
		require.ElementsMatch(t, []int{7, 8, 9}, tr.inorder())
		tr.insert(6)
		require.ElementsMatch(t, []int{6, 7, 8, 9}, tr.inorder())
		tr.insert(5)
		require.ElementsMatch(t, []int{5, 6, 7, 8, 9}, tr.inorder())
		tr.insert(4)
		require.ElementsMatch(t, []int{4, 5, 6, 7, 8, 9}, tr.inorder())
		tr.insert(3)
		require.ElementsMatch(t, []int{3, 4, 5, 6, 7, 8, 9}, tr.inorder())
		tr.insert(2)
		require.ElementsMatch(t, []int{2, 3, 4, 5, 6, 7, 8, 9}, tr.inorder())
		tr.insert(1)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tr.inorder())

		for _, k := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
			n := find(tr.root, k)
			bucketNode := &tr.buckets[n.pos.row][n.pos.col]
			require.Equal(t, n, bucketNode)
			pN, pBucketNode := unsafe.Pointer(n), unsafe.Pointer(bucketNode)
			require.Equal(t, pN, pBucketNode)
		}
	})

	t.Run("up", func(t *testing.T) {
		tr := newTree[int](0)

		tr.insert(1)
		require.ElementsMatch(t, []int{1}, tr.inorder())
		tr.insert(2)
		require.ElementsMatch(t, []int{1, 2}, tr.inorder())
		tr.insert(3)
		require.ElementsMatch(t, []int{1, 2, 3}, tr.inorder())
		tr.insert(4)
		require.ElementsMatch(t, []int{1, 2, 3, 4}, tr.inorder())
		tr.insert(5)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5}, tr.inorder())
		tr.insert(6)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6}, tr.inorder())
		tr.insert(7)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7}, tr.inorder())
		tr.insert(8)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, tr.inorder())
		tr.insert(9)
		require.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, tr.inorder())
	})

	t.Run("rand", func(t *testing.T) {
		for _, cnt := range []int{1, 10, 50, 100, 150, 200, 500} {
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
		empty := newTree[int](0)
		v := find(empty.root, 15)
		require.Nil(t, v)
	})

	t.Run("rand", func(t *testing.T) {
		cnt := 10
		vals, tr := testIntRandSliceToTree(t, cnt)

		v := find(tr.root, cnt+1)
		require.Nil(t, v)

		v = find(tr.root, 3)
		require.NotNilf(t, v, "vals %v, key %d", vals, 3)

		for _, k := range vals {
			n := find(tr.root, k)
			require.NotNilf(t, n, "vals %v, key %d", vals, k)
			require.Equal(t, k, n.data)
		}
	})
}

func Test_inorder_nonrec(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		empty := newTree[int](0)
		keys := empty.root.inorder(0)
		require.Equal(t, ([]int)(nil), keys)
	})

	t.Run("one", func(t *testing.T) {
		tr := newTree[int](0)
		tr.insert(1)
		keys := tr.root.inorder(1)
		require.Equal(t, []int{1}, keys)
	})

	t.Run("two", func(t *testing.T) {
		tr := newTree[int](0)
		tr.insert(2)
		tr.insert(1)
		keys := tr.root.inorder(1)
		require.Equal(t, []int{1, 2}, keys)
	})

	t.Run("three", func(t *testing.T) {
		tr := newTree[int](0)
		tr.insert(2)
		tr.insert(1)
		tr.insert(7)
		keys := tr.root.inorder(1)
		require.Equal(t, []int{1, 2, 7}, keys)
	})

	for _, cnt := range []int{5, 10, 25, 55, 100} {
		t.Run(strconv.Itoa(cnt), func(t *testing.T) {
			expKeys, tr := testIntRandSliceToTree(t, cnt)
			keys := tr.root.inorder(cnt)
			slices.Sort(expKeys)
			require.Equal(t, expKeys, keys)
		})
	}
}

func Test_iterator(t *testing.T) {
	for _, cnt := range []int{0, 5, 10, 25, 55, 100} {
		t.Run(strconv.Itoa(cnt), func(t *testing.T) {
			expKeys, tr := testIntRandSliceToTree(t, cnt)
			slices.Sort(expKeys)

			keys := make([]int, 0, len(expKeys))
			next := tr.root.iterator(100)
			for key := next(); key != nil; key = next() {
				keys = append(keys, *key)
			}

			require.Equal(t, expKeys, keys)
		})
	}
}

func testIntRandSliceToTree(t *testing.T, n int) ([]int, *tree[int]) {
	t.Helper()

	tr := newTree[int](0)

	vals := rand.Perm(n)
	for _, v := range vals {
		tr.insert(v)
	}

	return vals, tr
}
