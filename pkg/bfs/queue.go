package bfs

type node[T comparable] struct {
	Value    T
	Verified bool
	Next     *node[T]
}

type queue[T comparable] struct {
	Head *node[T]
	Tail *node[T]
}

func NewNode[T comparable](value T) *node[T] {
	n := node[T]{
		Value:    value,
		Verified: false,
	}

	return &n
}

func NewQueue[T comparable]() *queue[T] {
	q := queue[T]{}

	return &q
}

func (q *queue[T]) Enqueue(value T) {
	n := NewNode[T](value)

	if q.Head == nil {
		q.Head = n
		q.Tail = n
	} else {
		q.Tail.Next = n
		q.Tail = n
	}
}

func (q *queue[T]) Dequeue() *node[T] {
	if q.Head == nil {
		return nil
	}

	head := q.Head

	q.Head = head.Next

	return head
}

func (q *queue[T]) Values() []T {
	values := []T{}

	for curr := q.Head; curr != nil; curr = curr.Next {
		values = append(values, curr.Value)
	}

	return values
}
