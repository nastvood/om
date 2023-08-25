package om

/*func Test_find(t *testing.T) {
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

func testIntRandSliceToTree(t *testing.T, n int) ([]int, *node[int]) {
	t.Helper()

	var tree *node[int]

	vals := rand.Perm(n)
	for _, v := range vals {
		tree = insert(tree, v)
	}

	return vals, tree
}*/
