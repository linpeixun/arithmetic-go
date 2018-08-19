package sort

import (
	"fmt"
	"testing"
)

func Test_shell(t *testing.T) {
	arr := []int{9, 1, 2, 5, 7, 4, 8, 6, 3, 5}
	fmt.Println(shellSort(arr))
}
