package sort

import (
	"fmt"
	"testing"
)

func Test_buddle(t *testing.T) {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(bubbleSort(arr))

	fmt.Println(arr)
}
