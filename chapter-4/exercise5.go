package main

import "fmt"

func e5() {
	s := []string{"a", "a", "b", "c", "d", "d", "d", "e", "f", "g", "g"}
	fmt.Println(s)
	s = e5EDuplicate(s)
	fmt.Println(s)
}

// eliminates adjacent duplicates in s with in-place technique
func e5EDuplicate(s []string) []string {
	j := 0
	for i := 0; i < len(s); i++ {
		k := i
		for k < len(s) && s[i] == s[k] {
			k++
		}
		k--
		s[j] = s[k]
		j++
		i = k
	}
	return s[:j]
}
