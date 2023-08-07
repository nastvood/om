package oam

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Add(t *testing.T) {
	m := New[int, int]()
	m.Add(1, 2)
	m.Add(2, 4)
	m.Add(3, 6)

	require.Equal(t, 3, m.Len())
	m.Delete(3)

	m.Add(2, 4)
	require.Equal(t, 2, m.Len())

	m.Delete(1)
	require.Equal(t, 1, m.Len())

	v, ok := m.Get(2)
	require.True(t, ok)
	require.Equal(t, v, 4)

	v, ok = m.Get(1)
	require.False(t, ok)
	require.Equal(t, v, 0)

	require.ElementsMatch(t, []bool{false, true, false}, m.alive)
	require.ElementsMatch(t, []int{1, 2, 3}, m.keys)
	require.Equal(t, map[int]int{2: 1}, m.index)
	require.Equal(t, map[int]int{2: 4}, m.data)
}
