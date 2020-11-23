package alg

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i > -1; i-- {
		digits[i] = digits[i] + 1
		digits[i] = digits[i] % 10
		if digits[i] != 0 {
			return digits
		}

	}

	digits=append(digits, 0)
	copy(digits[1:],digits[:])
	digits[0]=1

	return digits
}
