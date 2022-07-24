package link2

type Link2[T any] struct {
	Node T
	Pre  *Link2[T]
	Next *Link2[T]
}

func (l *Link2[T]) PushBack(v T) {
	n := &Link2[T]{v, l, nil}
	if l.Next == nil {
		l.Next = n
		l.Pre = n
	} else {
		l.Next.Pre = n
		l.Next = n
	}
}

func (l *Link2[T]) PushFront(v T) {
	n := &Link2[T]{v, nil, l}
	if l.Pre == nil {
		l.Pre = n
		l.Next = n
	} else {
		l.Pre.Next = n
		l.Pre = n
	}
}

func (l *Link2[T]) IteratorNext(fn func(v *Link2[T]) bool) {
	if l.Next == nil {
		return
	}
	for n := l.Next; n != nil; n = n.Next {
		if !fn(n) {
			return
		}
	}
}

func (l *Link2[T]) IteratorFront(fn func(v *Link2[T]) bool) {
	if l.Pre == nil {
		return
	}
	for n := l.Pre; n != nil; n = n.Pre {
		if !fn(n) {
			return
		}
	}
}

func NewLink2[T any](v T) Link2[T] {
	return Link2[T]{
		Node: v,
		Pre:  nil,
		Next: nil,
	}
}

func (l *Link2[T]) DelNext(v Link2[T]) {
	if l.Next == nil {
		return
	}
	if l.Next == nil {
		return
	}
	if l.Next.Next != nil {
		l.Next.Next.Pre = l
	}
	l.Next = l.Next.Next
}

func (l *Link2[T]) FindVal(v T) *Link2[T] {
	var res *Link2[T]
	l.IteratorNext(func(n *Link2[T]) bool {
		// TODO Equal
		return true
	})
	return res
}
