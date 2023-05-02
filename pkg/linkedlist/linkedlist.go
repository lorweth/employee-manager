package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Root *Node[T]
}

func OfArray[T any](values []T) LinkedList[T] {
	ll := LinkedList[T]{}
	for _, v := range values {
		ll.Push(v)
	}

	return ll
}

func (l *LinkedList[T]) Push(value T) {
	if l.Root == nil {
		l.Root = new(Node[T])
		l.Root.Value = value
		return
	}
	l.Root.Next = new(Node[T])
	l.Root.Next.Value = value
}

func (l *LinkedList[T]) Find(cond func(value T) bool) *T {
	if l.Root != nil {
		cur := l.Root
		for cur != nil {
			if cond(cur.Value) {
				return &cur.Value
			}

			cur = cur.Next
		}
	}

	return nil
}

func (l *LinkedList[T]) Filter(cond func(value T) bool) LinkedList[T] {
	filtered := LinkedList[T]{}

	if l.Root != nil {
		cur := l.Root
		for cur != nil {
			if cond(cur.Value) {
				filtered.Push(cur.Value)
			}

			cur = cur.Next
		}
	}

	return filtered
}

func (l *LinkedList[T]) Walk(f func(value T)) {
	if l.Root != nil {
		cur := l.Root
		for cur != nil {
			f(cur.Value)

			cur = cur.Next
		}
	}
}
