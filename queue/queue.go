package queue

type Queue[T any] []T

func NewQueue[T any](cnt int) Queue[T] {
	return make(Queue[T], 0, cnt)
}

func (q *Queue[T]) PushBack(v T) {
	if len(*q) == cap(*q) {
		return
	}
	*q = append(*q, v)
}

func (q *Queue[T]) PushFront(v T) {
	if len(*q) == cap(*q) {
		return
	}
	*q = append([]T{v}, *q...)
}

func (q *Queue[T]) PopFront() *T {
	if len(*q) == 0 {
		return nil
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return &r
}

func (q *Queue[T]) PopBack() *T {
	if len(*q) == 0 {
		return nil
	}
	l := len(*q)
	r := (*q)[l-1]
	*q = (*q)[:l-1]
	return &r
}
