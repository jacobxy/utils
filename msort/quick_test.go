package msort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 2, 5, 1, 4, 1, 2, 3, 6, 5, 8, 74, 1, 2, 5, 4, 2, 6, 78, 46, 456, 23, 45, 64, 5646}
	t.Log(arr)
	QuickSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
