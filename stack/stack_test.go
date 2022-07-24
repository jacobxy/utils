package stack

import (
	"testing"
)

type TestStack struct {
	Val  int
	Name string
}

func TestNewStack(t *testing.T) {
	r := NewStack[TestStack]()
	r.Push(TestStack{Val: 1, Name: "test1"})
	r.Push(TestStack{Val: 2, Name: "test2"})
	res := r.Pop()
	t.Log(res)
	res = r.Pop()
	t.Log(res)
}
