package sort

func shellSort(arr []int) []int {
	increment := len(arr)
	for {
		increment = increment / 2
		for i := 0; i < increment; i++ {
			for j := i + increment; j < len(arr); j = j + increment {
				for k := j; k > i; k = k - increment {
					if arr[k] < arr[k-increment] {
						arr[k], arr[k-increment] = arr[k-increment], arr[k]
					} else {
						break
					}
				}
			}
		}
		if increment == 1 {
			break
		}
	}
	return arr
}
