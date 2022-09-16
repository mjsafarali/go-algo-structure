package queue

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestQueue(t *testing.T) {
	for title, fn := range map[string]func(t *testing.T){
		"get count of queue list": testLen,    // test: len functionality
		"get first item of queue": testPeek,   // test: peek functionality
		"remove item from queue":  testRemove, // test: remove functionality
	} {
		t.Run(title, func(t *testing.T) {
			fn(t)
		})
	}
}

func testLen(t *testing.T) {
	q := Queue{}
	var item interface{}

	require.Equal(t, 0, q.Len())

	q.Add(123)
	require.Equal(t, 1, q.Len())

	q.Add("test")
	require.Equal(t, 2, q.Len())

	q.Add(strings.NewReader(`{"a": "b", "c": {"en": "c"}}`))
	require.Equal(t, 3, q.Len())

	_, _ = q.Peek()
	require.Equal(t, 3, q.Len())

	item, _ = q.Remove()
	require.Equal(t, 2, q.Len())
	require.Equal(t, 123, item)
}

func testRemove(t *testing.T) {
	q := Queue{}
	var item interface{}

	q.Add(1)
	q.Add(2)
	q.Add(3)

	item, _ = q.Peek()
	require.Equal(t, 1, item)

	item, _ = q.Remove()
	require.Equal(t, 1, item)

	item, _ = q.Remove()
	require.Equal(t, 2, item)

	item, _ = q.Remove()
	require.Equal(t, 3, item)

	item, err := q.Remove()
	require.Error(t, ErrorEmpty, err)
}

func testPeek(t *testing.T) {
	q := Queue{}

	item, err := q.Peek()
	require.Error(t, ErrorEmpty, err)

	q.Add("test")
	item, err = q.Peek()
	require.NoError(t, nil)
	require.Equal(t, "test", item)

	q.Add(100)
	item, err = q.Peek()
	require.NoError(t, nil)
	require.Equal(t, "test", item)
	//queue: test,100
	item, err = q.Remove()
	require.NoError(t, nil)
	require.Equal(t, "test", item)

	item, err = q.Peek()
	require.NoError(t, nil)
	require.Equal(t, 100, item)
}
