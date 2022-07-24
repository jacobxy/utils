package find2mid

func Find(nums []int, val int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)>>1
		if val == nums[mid] {
			return mid
		} else if val < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
