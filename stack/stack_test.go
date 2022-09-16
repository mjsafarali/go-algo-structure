package stack

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestQueue(t *testing.T) {
	for title, fn := range map[string]func(t *testing.T){
		"get count of stack list": testLen,  // test: len functionality
		"get first item of stack": testPeek, // test: peek functionality
		"remove item from stack":  testPop,  // test: remove functionality
	} {
		t.Run(title, func(t *testing.T) {
			fn(t)
		})
	}
}

func testLen(t *testing.T) {
	s := Stack{}
	var item interface{}

	require.Equal(t, 0, s.Len())

	s.Push(strings.NewReader(`{"a": "b", "c": {"en": "c"}}`))
	require.Equal(t, 1, s.Len())

	s.Push("test")
	require.Equal(t, 2, s.Len())

	s.Push(123)
	require.Equal(t, 3, s.Len())

	_, _ = s.Peek()
	require.Equal(t, 3, s.Len())

	item, _ = s.Pop()
	require.Equal(t, 2, s.Len())
	require.Equal(t, 123, item)
}

func testPop(t *testing.T) {
	s := Stack{}
	var item interface{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	item, _ = s.Peek()
	require.Equal(t, 3, item)

	item, _ = s.Pop()
	require.Equal(t, 3, item)

	item, _ = s.Pop()
	require.Equal(t, 2, item)

	item, _ = s.Pop()
	require.Equal(t, 1, item)

	item, err := s.Pop()
	require.Error(t, ErrorEmpty, err)
}

func testPeek(t *testing.T) {
	s := Stack{}

	item, err := s.Peek()
	require.Error(t, ErrorEmpty, err)

	s.Push("test")
	item, err = s.Peek()
	require.NoError(t, err)
	require.Equal(t, "test", item)

	s.Push(100)
	item, err = s.Peek()
	require.NoError(t, err)
	require.Equal(t, 100, item)

	item, err = s.Pop()
	require.NoError(t, err)
	require.Equal(t, 100, item)

	item, err = s.Peek()
	require.NoError(t, err)
	require.Equal(t, "test", item)
}
