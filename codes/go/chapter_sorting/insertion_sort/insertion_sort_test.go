// File: insertion_sort_test.go
// Created Time: 2022-12-12
// Author: msk397 (machangxinq@gmail.com)

package insertion_sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{4, 1, 3, 1, 5, 2}
	insertionSort(nums)
	fmt.Println("插入排序完成后 nums =", nums)
}
