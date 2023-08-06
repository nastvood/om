package oam

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_iter_next(t *testing.T) {
	m := New[string, int](10, false)
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
	m.Del("4")
	m.Add("5", 5)
	m.Del("5")
	m.Add("6", 6)
	m.Add("5", 5)
	m.Del("5")

	k, ok = iter.Next()
	require.True(t, ok)
	require.Equal(t, "6", k)

	k, ok = iter.Next()
	require.False(t, ok)
	require.Equal(t, "", k)
}
