package main

import "fmt"

func e11() {
	fmt.Println(e11Comma("-12345678101112131.423617"))
}

// comma (non-recursive version) inserts commas in a non-negative decimal integer string.
func e11Comma(s string) string {
	var sign bool
	if s[0] == '-' {
		sign = true
		s = s[1:]
	}
	i := len(s)
	for j, c := range s {
		if c == '.' {
			i = j
		}
	}
	for i -= 3; i > 0; i -= 3 {
		s = s[:i] + "," + s[i:]
	}
	if sign {
		s = "-" + s
	}
	return s
}
