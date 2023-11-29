package bfs

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run(
		"Enqueue() should add the node on head and tail if queue is empty",
		func(t *testing.T) {
			q := NewQueue[string]()

			q.Enqueue("Alice")

			if q.Head.Value != "Alice" || q.Tail.Value != "Alice" {
				t.Errorf(
					"expected head and tail values to be %s, got head: %s, tail: %s",
					"Alice",
					q.Head.Value,
					q.Tail.Value,
				)
			}
		},
	)

	t.Run("Enqueue() should add node behind the node in front", func(t *testing.T) {
		q := NewQueue[string]()

		q.Enqueue("Bob")
		q.Enqueue("Charlie")

		if q.Head.Value != "Bob" || q.Tail.Value != "Charlie" {
			t.Errorf(
				"expected head to be %s and tail to be %s, got head: %s, tail: %s",
				"Bob",
				"Charlie",
				q.Head.Value,
				q.Tail.Value,
			)
		}
	})

	t.Run(
		"Dequeue() should remove current first node in the queue and return nil if queue is empty",
		func(t *testing.T) {
			q := NewQueue[string]()

			dequeued := q.Dequeue()

			if dequeued != nil {
				t.Errorf("expected dequeued node to be nil, got %+v", dequeued)
			}

			q.Enqueue("Jhon")
			q.Enqueue("Snow")

			dequeued = q.Dequeue()

			if dequeued.Value != "Jhon" {
				t.Errorf("expected dequeued node value to be %s, got %s", "Jhon", dequeued.Value)
			}

			if q.Head.Value != "Snow" {
				t.Errorf("expected head node value to be %s, got %s", "Snow", q.Head.Value)
			}
		},
	)

	t.Run("Values() should return slice of queue values", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		values := q.Values()

		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(values, expected) {
			t.Errorf("expected values to be %v, got %v", expected, values)
		}
	})

	t.Run("EnqueueMany() should enqueue many items", func(t *testing.T) {
		q := NewQueue[int]()

		items := []int{1, 2, 3}
		empty := []int{}

		q.EnqueueMany(empty)

		if q.Head != nil {
			t.Errorf("expect %v - got %+v", nil, q.Head)
		}

		q.EnqueueMany(items)

		if q.Head == nil || q.Head.Value != items[0] {
			t.Errorf("expect value %d - got %v", items[0], q.Head)
		}

		if q.Tail == nil || q.Tail.Value != items[2] {
			t.Errorf("expect value %d - got %v", items[2], q.Tail)
		}
	})
}
