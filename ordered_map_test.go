package om

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"golang.org/x/exp/slices"
)

func Test_M(t *testing.T) {
	m := New[int, string]()

	v, ok := m.Get(1)
	require.False(t, ok)
	require.Equal(t, "", v)

	cnt := 25
	keys := rand.Perm(cnt)
	for _, k := range keys {
		m.Add(k, strconv.Itoa(k))
	}

	for _, k := range keys {
		v, ok = m.Get(k)
		require.True(t, ok)
		require.Equal(t, strconv.Itoa(k), v)
	}

	sortedKeys := make([]int, len(keys))
	copy(sortedKeys, keys)
	slices.Sort(sortedKeys)
	require.Equal(t, sortedKeys, m.Keys())
}
