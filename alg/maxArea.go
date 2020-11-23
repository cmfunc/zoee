package alg

// maxArea 盛最多水的容器
func maxArea(height []int) int {
	var max int
	for i, val_i := range height {
		for j, val_j := range height {
			var x, y int
			if i < j {
				x = j - i
			} else {
				x = i - j
			}
			if val_j < val_i {
				y = val_j
			} else {
				y = val_i
			}

			if max < x*y {
				max = x * y
			}
		}
	}
	return max
}
