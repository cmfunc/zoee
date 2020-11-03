package alg

// TwoSum 两数字之和
func TwoSum(nums []int, target int) []int {
	tmpMap := make(map[int]int, 0)
	for i, num := range nums {
		if j, ok := tmpMap[num]; ok && i!=j {
			return []int{j, i}
		}
		tmpMap[target-num] = i
	}
	return []int{}
}
