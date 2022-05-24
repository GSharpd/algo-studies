package main

import (
	"fmt"
	"sort"
)

func main() {
	arr1 := []int{0, 1, 3, 5, 7, 9}
	arr2 := []int{2, 4, 6, 8, 10}
	fmt.Println(mergeSortedArray(arr1, arr2))
}

func mergeSortedArray(a, b []int) (result []int) {
	res := make([]int, 0)
	for _, v := range a {
		res = append(res, v)
	}

	for _, v := range b {
		res = append(res, v)
	}

	sort.Ints(res)

	return res
}
