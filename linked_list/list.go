package singlylinkedlist

// Node represents a node in singly linked list.
type Node struct {
	Value int
	Next  *Node
}

// List implements a singly linked list.
type List struct {
	first *Node
}

func newList(values ...int) *List {
	if len(values) < 1 {
		return &List{}
	}

	temp := &Node{}
	current := temp
	for _, v := range values {
		current.Next = &Node{Value: v}
		current = current.Next
	}

	return &List{first: temp.Next}
}

// Append adds an item to the list.
func (l *List) Append(value int) {
	newNode := &Node{Value: value}

	if l.first == nil {
		l.first = newNode
		return
	}

	current := l.first
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// Values returns the sequence of items in the list.
func (l *List) Values() []int {
	values := []int{}

	current := l.first
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}

	return values
}
