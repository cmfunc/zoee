package alg

// BubbleSort 冒泡
func BubbleSort(arrayOrigin []int) []int {

	for i, _ := range arrayOrigin {

		for j,_ := range arrayOrigin[:len(arrayOrigin)-i-1] {

			if arrayOrigin[j] > arrayOrigin[j+1] {
				arrayOrigin[j], arrayOrigin[j+1] = arrayOrigin[j+1], arrayOrigin[j]
			}

		}

	}

	return arrayOrigin
}
