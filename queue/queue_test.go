package queue

import (
	"testing"
)

type data struct {
	Val  int
	Name string
}

func TestQueue(t *testing.T) {
	NewQueue[data](10)
}
