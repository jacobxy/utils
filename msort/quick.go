package msort

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	i, j := left, right
	pivot := arr[(left+right)/2]

	for i <= j {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	QuickSort(arr, left, j)
	QuickSort(arr, i, right)
}
