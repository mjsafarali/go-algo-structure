package singlylinkedlist

// SinglyLinkedList implements a list of items.
type SinglyLinkedList interface {
	Append(value int)
	Values() []int
}

// NewSinglyLinkedList return a new list.
func NewSinglyLinkedList(values ...int) SinglyLinkedList {
	return newList(values...)
}
