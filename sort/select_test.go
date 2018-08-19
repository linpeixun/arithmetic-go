package sort

import (
	"fmt"
	"testing"
)

func Test_select(t *testing.T) {
	arr := []int{70, 30, 40, 10, 80, 20, 90, 100, 75, 60, 45}
	fmt.Println(selectSort(arr))
}
