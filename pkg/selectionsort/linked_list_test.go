package selectionsort

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run(
		"the first item to be inserted as the head node must also be the tail node",
		func(t *testing.T) {
			l := NewLinkedList[string]()

			l.AddAtTheBeginning("Gabriel")

			if l.Head.Value != l.Tail.Value {
				t.Errorf("expect %s - receive %s", l.Head.Value, l.Tail.Value)
			}
		},
	)

	t.Run(
		"the first item to be inserted as the tail node must also be the head node if list is empty",
		func(t *testing.T) {
			l := NewLinkedList[string]()

			l.AddAtTheEnd("Mariana")

			if l.Head.Value != l.Tail.Value {
				t.Errorf("expect %s - receive %s", l.Head.Value, l.Tail.Value)
			}
		},
	)

	t.Run("should add node on the tail", func(t *testing.T) {
		l := NewLinkedList[string]()

		l.AddAtTheBeginning("Gabriel")

		l.AddAtTheEnd("Mariana")

		if l.Tail.Value != "Mariana" {
			t.Errorf("expect %s - receive %s", "Mariana", l.Tail.Value)
		}

		if l.Head.Next.Value != "Mariana" {
			t.Errorf("expect %s - receive %s", "Mariana", l.Head.Next.Value)
		}

		if l.Tail.Value != "Mariana" {
			t.Errorf("expect %s - receive %s", "Mariana", l.Tail.Value)
		}
	})

	t.Run("should add node on the head", func(t *testing.T) {
		l := NewLinkedList[string]()

		l.AddAtTheBeginning("Initial")

		l.AddAtTheBeginning("New Initial")

		if l.Head.Value != "New Initial" {
			t.Errorf("expect %s - receive %s", "New Initial", l.Tail.Value)
		}

		if l.Head.Next.Value != "Initial" {
			t.Errorf("expect %s - receive %s", "Initial", l.Head.Next.Value)
		}

		if l.Tail.Value != "Initial" {
			t.Errorf("expect %s - receive %s", "Initial", l.Tail.Value)
		}
	})

	t.Run("should remove node on the tail", func(t *testing.T) {
		l := NewLinkedList[string]()

		l.AddAtTheBeginning("First")
		l.AddAtTheEnd("Secondary")
		l.AddAtTheEnd("Third")

		if l.Tail.Value != "Third" {
			t.Errorf("expect %s - receive %s", "Third", l.Tail.Value)
		}

		removed := l.PopTail()

		if removed.Value != "Third" {
			t.Errorf("expect %s - receive %s", "Third", removed.Value)
		}

		if l.Tail.Next != nil {
			t.Errorf("expect %v - receive %+v", nil, l.Tail.Next)
		}

		if l.Tail.Value != "Secondary" {
			t.Errorf("expect %s - receive %s", "Secondary", l.Tail.Value)
		}
	})

	t.Run("should remove node on the head", func(t *testing.T) {
		l := NewLinkedList[string]()

		l.AddAtTheBeginning("First")
		l.AddAtTheEnd("Secondary")

		removed := l.PopHead()

		if removed.Value != "First" {
			t.Errorf("expect %s - receive %s", "First", removed.Value)
		}

		if l.Head.Value != "Secondary" {
			t.Errorf("expect %s - receive %s", "Secondary", l.Head.Value)
		}

		if l.Head != l.Tail {
			t.Errorf("expect %+v - receive %+v", l.Tail, l.Head)
		}

		if l.Head.Next != nil {
			t.Errorf("expect %v - receive %v", nil, l.Head.Next)
		}
	})

	t.Run("should display full list items value", func(t *testing.T) {
		l := NewLinkedList[int]()

		l.AddAtTheBeginning(1)
		l.AddAtTheEnd(2)
		l.AddAtTheEnd(3)
		l.AddAtTheEnd(4)

		listValues := l.Display()

		expect := []int{1, 2, 3, 4}

		if !reflect.DeepEqual(listValues, expect) {
			t.Errorf("expect %v - receive %v - tail %+v", expect, listValues, *l.Tail)
		}
	})

	t.Run("should find item on list by value", func(t *testing.T) {
		l := NewLinkedList[int]()

		l.AddAtTheBeginning(1)
		l.AddAtTheEnd(2)
		l.AddAtTheEnd(3)
		l.AddAtTheEnd(4)
		l.AddAtTheEnd(5)

		nodeFinded := l.Find(4)

		if nodeFinded == nil || nodeFinded.Value != 4 {
			t.Errorf("expect %+v - receive %v", *NewNode[int](4), nodeFinded)
		}
	})
}
