package oam

import (
	"strconv"
	"testing"

	"github.com/nastvood/om/conf"
	"github.com/stretchr/testify/require"
)

func Test_iter_next(t *testing.T) {
	m := New[string, int](conf.WithCapacity(10))
	m.Add("1", 1)
	m.Add("2", 2)

	iter := m.Iterator()
	k, ok := iter.Next()
	require.True(t, ok)
	require.Equal(t, "1", k)

	k, ok = iter.Next()
	require.True(t, ok)
	require.Equal(t, "2", k)

	k, ok = iter.Next()
	require.False(t, ok)
	require.Equal(t, "", k)

	m.Add("3", 3)
	k, ok = iter.Next()
	require.True(t, ok)
	require.Equal(t, "3", k)

	m.Add("4", 4)
	m.Delete("4")
	m.Add("5", 5)
	m.Delete("5")
	m.Add("6", 6)
	m.Add("5", 5)
	m.Delete("5")

	k, ok = iter.Next()
	require.True(t, ok)
	require.Equal(t, "6", k)

	k, ok = iter.Next()
	require.False(t, ok)
	require.Equal(t, "", k)
}

func Test_iter_prev(t *testing.T) {
	m := New[string, int](conf.WithCapacity(10), conf.WithoutConcurrency())
	m.Add("1", 1)
	m.Add("2", 2)
	m.Add("3", 4)
	m.Add("4", 4)
	m.Add("1", 1)
	m.Add("2", 2)
	m.Delete("4")
	m.Delete("3")
	m.Add("5", 5)
	m.Add("6", 6)

	iter := m.Iterator()
	k, ok := iter.Prev()
	require.False(t, ok)
	require.Equal(t, "", k)

	iter.End()

	for _, exp := range []int{6, 5, 2, 1} {
		k, ok = iter.Prev()
		require.True(t, ok)
		require.Equal(t, strconv.Itoa(exp), k)
	}

	iter.End()
	k, ok = iter.Prev()
	require.True(t, ok)
	require.Equal(t, "6", k)

	iter.Begin()
	k, ok = iter.Prev()
	require.False(t, ok)
	require.Equal(t, "", k)

	for _, exp := range []int{1, 2, 5} {
		k, ok = iter.Next()
		require.True(t, ok)
		require.Equal(t, strconv.Itoa(exp), k)
	}

	k, ok = iter.Prev()
	require.True(t, ok)
	require.Equal(t, "2", k)

	k, ok = iter.Next()
	require.True(t, ok)
	require.Equal(t, "5", k)
}
