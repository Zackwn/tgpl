package main

import "fmt"

func e4() {
	s := []int{1, 2, 3}
	n := 2
	fmt.Println(s)
	fmt.Println(e4Rotate(s, n))
}

// rotates s left in one pass
func e4Rotate(s []int, n int) []int {
	n = n % len(s)
	if n > 0 && n < len(s) {
		temp := make([]int, n)
		copy(temp, s[:n])
		copy(s[:n], s[n:])
		copy(s[len(s)-n:], temp)
	}
	return s
}
