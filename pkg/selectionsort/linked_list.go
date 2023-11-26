package selectionsort

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

type linkedList[T comparable] struct {
	Head *node[T]
	Tail *node[T]
}

func NewNode[T comparable](value T) *node[T] {
	n := node[T]{
		Value: value,
		Next:  nil,
	}

	return &n
}

func NewLinkedList[T comparable]() *linkedList[T] {
	l := linkedList[T]{}

	return &l
}

func (l *linkedList[T]) AddAtTheBeginning(value T) {
	node := NewNode[T](value)

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		node.Next = l.Head

		l.Head = node
	}
}

func (l *linkedList[T]) AddAtTheEnd(value T) {
	node := NewNode[T](value)

	if l.Tail == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *linkedList[T]) Next() *node[T] {
	return l.Head.Next
}

func (l *linkedList[T]) PopHead() *node[T] {
	if l.Head == nil {
		return nil
	}

	oldHead := l.Head
	l.Head = oldHead.Next

	return oldHead
}

func (l *linkedList[T]) PopTail() *node[T] {
	oldTail := l.Tail

	if l.Head == nil {
		return nil
	}

	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil

		return oldTail
	}

	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Next == oldTail {
			curr.Next = nil

			l.Tail = curr
		}
	}

	return oldTail
}

func (l *linkedList[T]) Display() []T {
	values := []T{}

	for curr := l.Head; curr != nil; curr = curr.Next {
		values = append(values, curr.Value)
	}

	return values
}

func (l *linkedList[T]) Find(value T) *node[T] {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if curr.Value == value {
			return curr
		}
	}

	return nil
}
