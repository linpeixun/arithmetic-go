package sort

import (
	"fmt"
	"testing"
)

func Test_insert(t *testing.T) {
	arr := []int{3, 6, 4, 2, 11, 10, 5}
	fmt.Println(insertSort(arr))
}
