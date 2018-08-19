package sort

import (
	"fmt"
	"testing"
)

func Test_heap(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(heapSort(arr))
}
