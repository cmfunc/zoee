package alg

// validMountainArray 山脉数组
func validMountainArray(A []int) bool {

	if len(A) < 3 {
		return false
	}

	var i, j int
	var max_i, max_j int
	for i = 0; i < len(A)-1; i++ {
        if A[i]==A[i+1]{
            return false
        }
		if A[i] > A[i+1] {
			max_i = A[i]
			break
		}
	}
	for j = len(A) - 1; j > 0; j-- {
        if A[j]==A[j-1]{
            return false
        }
		if A[j] > A[j-1] {
			max_j = A[j]
			break
		}
	}
	return i == j && max_i == max_j
}
