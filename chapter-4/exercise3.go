package main

import "fmt"

func e3() {
	var arr [e3Size]int
	for i := 0; i < e3Size; i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	arr = e3Reverse(&arr)
	fmt.Println(arr)
}

const e3Size = 5

// reverse with array pointers
func e3Reverse(arr *[e3Size]int) [e3Size]int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return *arr
}
