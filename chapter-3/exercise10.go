package main

import "fmt"

func e10() {
	fmt.Println(e10Comma("12345678101112131"))
}

// comma (non-recursive version) inserts commas in a non-negative decimal integer string.
func e10Comma(s string) string {
	n := len(s)
	for i := n - 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	return s
}
