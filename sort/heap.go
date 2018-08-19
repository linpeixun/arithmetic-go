package sort

func heapSort(arr []int) []int {
	length := len(arr)

	//构建大堆
	for i := length/2 - 1; i >= 0; i-- {
		//从第一个非叶子结点从下至上，从右至左调整结构
		adjustHeap(arr, i, length)
	}
	//2.调整堆结构+交换堆顶元素与末尾元素
	for j := length - 1; j > 0; j-- {
		arr[0], arr[j] = arr[j], arr[0] //将堆顶元素与末尾元素进行交换
		adjustHeap(arr, 0, j)           //重新对堆进行调整
	}

	return arr
}

func adjustHeap(arr []int, i int, length int) {
	temp := arr[i]
	for k := i*2 + 1; k < length; k = k*1 + 1 {
		if k+1 < length && arr[k] < arr[k+1] {
			k++
		}

		if arr[k] > temp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}

		arr[i] = temp
	}
}
