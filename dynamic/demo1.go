package dynamic

var foot = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// leetcode 70. Climbing Stairs
func ClimStairs(n int) int {
	if n <= 1 {
		return n
	}
	res := make([]int, n)
	m := len(foot)

	res[0] = 1
	res[1] = 2

	for i := 2; i < n; i++ {
		for j := 0; j < m; j++ {
			if i > foot[j] {
				break
			}
			res[i] += res[i-foot[j]]
		}
	}

	return res[n-1]
}
